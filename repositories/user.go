// Package repositories
// @Author shaofan
// @Date 2022/5/13
package repositories

import "douyin/entity/po"

// User 用户持久层接口
type User interface {
	// Insert 			插入一条数据
	// user  			用户数据
	Insert(user *po.User) (int, error)

	// QueryById 		通过id查询
	// userId 			用户id
	// @return 			用户信息
	QueryById(userId int) (*po.User, error)

	// QueryBatchIds 	id批量查询
	// userIds 			用户id集
	// @return			用户集
	QueryBatchIds(userIds *[]int) (*[]po.User, error)
	// QueryByName		通过Name查询
	// userName			用户Name
	// @return 			用户信息
	QueryByName(userName string) (*po.User, error)
	//UpdateFollowById  通过id修改Follow
	//id				用户id
	//t					变化的数量
	UpdateFollowById(id int, t int) error
	//UpdateFollowerById  通过id修改Follower
	//id				用户id
	//t					变化的数量
	UpdateFollowerById(id int, t int) error
}
