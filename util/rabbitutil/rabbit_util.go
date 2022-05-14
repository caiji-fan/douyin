// Package rabbitutil
// @Author shaofan
// @Date 2022/5/13
// @DESC rabbitmq连接初始化与工具
package rabbitutil

func Init() {

}

// ChangeFollowNum 		修改用户粉丝数和关注数
// userId 				发起关注或取关的用户id
// toUserId 			收到关注或取关的用户id
// isFollow 			是否是关注请求
func ChangeFollowNum(userId int, toUserId int, isFollow bool) error {
	return nil
}

// UploadVideo 			上传视频文件
// filePath 			视频文件路径
func UploadVideo(filePath string) error {
	return nil
}

// FeedVideo 			投放视频到用户feed流
// videoId 				视频id
func FeedVideo(videoId int) error {
	return nil
}
