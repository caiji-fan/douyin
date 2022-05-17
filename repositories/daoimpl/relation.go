// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/entity/po"
	"douyin/repositories"
	"sync"
)

type Relation struct {
}

func (r Relation) Insert(follow *po.Follow) error {
	//TODO implement me
	panic("implement me")
}

func (r Relation) DeleteByCondition(follow *po.Follow) error {
	//TODO implement me
	panic("implement me")
}

func (r Relation) QueryFollowIdByFansId(fansId int) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (r Relation) QueryFansIdByFollowId(followId int) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

var (
	relation     repositories.Relation
	relationOnce sync.Once
)

func NewRelationDaoInstance() repositories.Relation {
	relationOnce.Do(func() {
		relation = Relation{}
	})
	return relation
}
