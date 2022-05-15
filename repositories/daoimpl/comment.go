// Package daoimpl
// @Author shaofan
// @Date 2022/5/13
package daoimpl

import (
	"douyin/repositories"
	"sync"
)

type Comment struct {
}

var (
	comment     repositories.Comment
	commentOnce sync.Once
)
