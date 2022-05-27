// Package serviceimpl
// @Author shaofan
// @Date 2022/5/13
package serviceimpl

import (
	"douyin/config"
	"douyin/entity/bo"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"douyin/service"
	"douyin/util/entityutil"
	"douyin/util/obsutil"
	"douyin/util/rabbitutil"
	"douyin/util/redisutil"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type Video struct {
}

func (v Video) Feed(userId int, isLogin bool, latestTime int64) ([]bo.Video, int64, error) {
	videoDao := daoimpl.NewVideoDaoInstance()
	// 视频集合
	var videos []po.Video
	// 从redis中取出的视频id集合
	var videoIds []int
	// 收件箱的feed
	var inbox []bo.Feed
	// 发件箱的feed
	var outbox []bo.Feed
	// 可能出现的错误
	var err error
	// 从feed数据库中删除数据的事务
	var tx *gorm.DB = nil
	// 如果用户已登录，查询redis中的feed流
	if isLogin {
		wait := sync.WaitGroup{}
		wait.Add(2)
		go func() {
			defer wait.Done()
			inbox, tx, err = fromInbox(userId)
		}()
		go func() {
			defer wait.Done()
			outbox, err = fromOutbox(userId)
		}()
		wait.Wait()
		if err != nil {
			if tx != nil {
				tx.Rollback()
			}
			return nil, 0, err
		}
		videoIds, err = mergeBox(&inbox, &outbox, latestTime)
		if err != nil {
			if tx != nil {
				tx.Rollback()
			}
			return nil, 0, err
		}
	}
	// 得到视频id集合，从数据库中查询视频数据
	videos, err = videoDao.QueryBatchIds(&videoIds, config.Config.Service.PageSize)
	if err != nil {
		tx.Rollback()
		return nil, 0, err
	}
	// 获取查询视频的时间条件，如果从redis从查到了数据，则为数据最后一条的时间，否则为前端传来的时间
	var timeCondition time.Time
	if len(videos) == 0 {
		timeCondition = time.UnixMilli(latestTime)
	} else {
		timeCondition = videos[len(videos)-1].CreateTime
	}
	// 如果用户未登录，或者feed流中的视频不足一页，通过时间条件在数据库中查询补足一页
	if len(videos) < config.Config.Service.PageSize {
		videos1, err := videoDao.QueryByLatestTimeDESC(timeCondition, config.Config.Service.PageSize-len(videos))
		if err != nil {
			if tx != nil {
				tx.Rollback()
			}
			return nil, 0, err
		}
		//拼接到原来数据后面
		videos = append(videos, *videos1...)
	}
	// 转换视频bo
	var videoBOS = make([]bo.Video, len(videos))
	err = entityutil.GetVideoBOS(&videos, &videoBOS)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return nil, 0, err
	}
	// 取出部分后，重新存入数据
	err = clearInbox(&inbox, userId)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		return nil, 0, err
	}
	return videoBOS, videos[len(videos)-1].CreateTime.UnixMilli(), nil
}

// Publish check token then save upload file to public directory
func (v Video) Publish(c *gin.Context, video *multipart.FileHeader, cover *multipart.FileHeader, userId int, title string) error {
	// 视频、封面本地保存
	videoName := obsutil.ParseFileName(filepath.Base(video.Filename))
	videoSaveFile := filepath.Join(config.Config.Service.VideoTempDir, videoName)
	if err := c.SaveUploadedFile(video, videoSaveFile); err != nil {
		return err
	}

	coverName := obsutil.ParseFileName(filepath.Base(cover.Filename))
	coverSaveFile := filepath.Join(config.Config.Service.CoverTempDir, coverName)
	if err := c.SaveUploadedFile(video, coverSaveFile); err != nil {
		return err
	}
	var videoDB = daoimpl.NewVideoDaoInstance()
	var tx = videoDB.Begin()
	videoPo := po.Video{
		PlayUrl:       videoSaveFile,
		CoverUrl:      coverSaveFile,
		FavoriteCount: 0,
		CommentCount:  0,
		AuthorId:      userId,
		Title:         title,
	}
	err := videoDB.Insert(tx, &videoPo, true)
	if err != nil {
		tx.Rollback()
		return err
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	// 消息队列异步上传视频， 并更新视频、封面的URL信息， 删除本地视频
	go func() {
		defer wg.Done()
		err = rabbitutil.UploadVideo(videoPo.ID)
	}()
	// 消息队列异步将视频加入feed流,正确响应
	go func() {
		err = rabbitutil.FeedVideo(videoPo.ID)
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (v Video) VideoList(userId int) ([]bo.Video, error) {
	// 查询数据库获取投稿列表
	poVideoList, err := daoimpl.NewVideoDaoInstance().QueryVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	var boVideoList = make([]bo.Video, len(*poVideoList))
	// po列表转bo
	err = entityutil.GetVideoBOS(poVideoList, &boVideoList)
	if err != nil {
		return nil, err
	}
	return boVideoList, nil
}

// 从自己的收件箱获取
func fromInbox(userId int) ([]bo.Feed, *gorm.DB, error) {
	var tx *gorm.DB
	var feedBOS = make([]bo.Feed, 0)
	err := redisutil.ZGet(config.Config.Redis.Key.Inbox+strconv.Itoa(userId), &feedBOS)
	if err != nil {
		return nil, nil, err
	}
	// 如果收件箱不存在，则从数据库中查询
	if len(feedBOS) == 0 {
		feedDao := daoimpl.NewFeedDaoInstance()
		tx = feedDao.Begin()
		feedPOS, err := feedDao.QueryByCondition(&po.Feed{UserId: userId})
		if err != nil {
			return nil, nil, err
		}
		err = feedDao.DeleteByCondition(&po.Feed{UserId: userId}, tx, true)
		if err != nil {
			return nil, nil, err
		}
		entityutil.GetFeedBOS(&feedPOS, &feedBOS)
	}
	return feedBOS, tx, nil
}

// 从大v的发件箱获取Feed流
// userId 			当前用户id
// @return			当前用户所有关注的大v的发件箱排序后的切片
func fromOutbox(userId int) ([]bo.Feed, error) {
	userDao := daoimpl.NewUserDaoInstance()
	follows, err := userDao.QueryFollows(userId)
	if err != nil {
		return nil, err
	}
	var feeds = make([]bo.Feed, 5)
	var length = 0
	for _, user := range *follows {
		if user.FollowerCount >= config.Config.Service.BigVNum {
			var feed = make([]bo.Feed, 5)
			err := redisutil.ZGet(config.Config.Redis.Key.Outbox+strconv.Itoa(user.ID), &feed)
			if err != nil {
				return nil, err
			}
			feeds, err = mergeFeeds(&feeds, &feed)
			if err != nil {
				return nil, err
			}
			length += len(feed)
			if length > config.Config.Service.PageSize {
				return feeds, nil
			}
		}
	}
	return feeds, nil
}

// 合并发件箱和收件箱的feed流，得到查询的id集合
// inbox 用户收件箱
// outbox 大v发件箱
// @return id集合
func mergeBox(inbox *[]bo.Feed, outbox *[]bo.Feed, latestTime int64) ([]int, error) {
	var videoIds = make([]int, config.Config.Service.PageSize)
	var inboxIndex, outboxIndex = 0, 0

	for inboxIndex < len(*inbox) && outboxIndex < len(*outbox) {
		inboxTime := (*inbox)[inboxIndex].CreateTime
		outboxTime := (*outbox)[outboxIndex].CreateTime
		if inboxTime.Before(outboxTime) {
			videoIds = append(videoIds, (*inbox)[inboxIndex].VideoId)
			inboxIndex++
		} else {
			// 只对发件箱中上一次latestTime之前的视频进行投送
			if outboxTime.UnixMilli() < latestTime {
				videoIds = append(videoIds, (*outbox)[outboxIndex].VideoId)
			}
			outboxIndex++
		}
		if len(videoIds) == config.Config.Service.PageSize {
			*inbox = (*inbox)[inboxIndex:]
			return videoIds, nil
		}
	}
	for inboxIndex < len(*inbox) {
		videoIds = append(videoIds, (*inbox)[inboxIndex].VideoId)
		inboxIndex++
		if len(videoIds) == config.Config.Service.PageSize {
			// 去掉没有选中的信息，为后面加锁管理redis中的收件箱做准备
			*inbox = (*inbox)[:inboxIndex+1]
			return videoIds, nil
		}
	}
	// 去掉没有选中的信息，为后面加锁管理redis中的收件箱做准备
	*inbox = (*inbox)[:inboxIndex+1]
	for outboxIndex < len(*outbox) {
		outboxTime := (*outbox)[outboxIndex].CreateTime
		// 只对发件箱中上一次latestTime之前的视频进行投送
		if outboxTime.UnixMilli() < latestTime {
			videoIds = append(videoIds, (*outbox)[outboxIndex].VideoId)
		}
		outboxIndex++
		if len(videoIds) == config.Config.Service.PageSize {
			return videoIds, nil
		}
	}
	return videoIds, nil
}

// 按顺序合并两个feed切片
func mergeFeeds(feed1 *[]bo.Feed, feed2 *[]bo.Feed) ([]bo.Feed, error) {
	var index1, index2 = 0, 0
	var feeds = make([]bo.Feed, len(*feed1)+len(*feed2))

	for index1 < len(*feed1) && index2 < len(*feed2) {
		time1 := (*feed1)[index1].CreateTime
		time2 := (*feed2)[index1].CreateTime
		if time1.Before(time2) {
			feeds = append(feeds, (*feed1)[index1])
			index1++
		} else {
			feeds = append(feeds, (*feed2)[index2])
			index2++
		}
	}
	for index1 < len(*feed1) {
		feeds = append(feeds, (*feed1)[index1])
		index1++
	}
	for index2 < len(*feed2) {
		feeds = append(feeds, (*feed2)[index2])
		index2++
	}
	return feeds, nil
}

// 清理收件箱，用户查看一次收件箱后，将收件箱中已经查看过的视频清除
// trash 已经查看过的feed对象
func clearInbox(trash *[]bo.Feed, userId int) error {
	//todo 加锁,防止清理过程中的新增数据造成影响

	// 获取redis中的数据
	var feeds []bo.Feed
	err := redisutil.ZGet(config.Config.Redis.Key.Outbox+strconv.Itoa(userId), &feeds)
	if err != nil {
		return err
	}
	// 获取垃圾中的视频id映射
	var mapper = make(map[int]int, len(*trash))
	for _, feed := range *trash {
		mapper[feed.VideoId] = 1
	}
	// 新切片从原切片中取非垃圾部分
	var newFeeds = make([]bo.Feed, 0)
	for _, feed := range feeds {
		if mapper[feed.VideoId] != 1 {
			newFeeds = append(newFeeds, feed)
		}
	}
	var value = make([]redis.Z, len(newFeeds))
	for i, v := range newFeeds {
		value[i] = redis.Z{
			Score:  float64(v.CreateTime.UnixMilli()),
			Member: v,
		}
	}
	err = redisutil.ZSetWithExpireTime(config.Config.Redis.Key.Outbox+strconv.Itoa(userId),
		value,
		config.OutboxExpireTime,
		false,
		nil)
	if err != nil {
		return err
	}
	return nil
}

var (
	video     service.Video
	videoOnce sync.Once
)

func NewVideoServiceInstance() service.Video {
	videoOnce.Do(func() {
		video = Video{}
	})
	return video
}
