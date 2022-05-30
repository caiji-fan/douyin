package entityutil

import (
	"douyin/config"
	"douyin/entity/bo"
	"douyin/entity/po"
	"douyin/repositories/daoimpl"
	"fmt"
	"testing"
	"time"
)

func TestMain(t *testing.M) {
	config.Init()
	daoimpl.Init()
	t.Run()
}
func TestTest(t *testing.T) {
	var cu = make(map[int]int, 5)
	cu[8] = 0
	cu[2] = 29
	cu[52] = 1
	cu[111] = 11
	cu[12] = 12
	//for i:=0;i<5;i++{
	//
	//}
	fmt.Println(cu)
}
func TestGetCommentBOS(t *testing.T) {
	var coms []po.Comment = []po.Comment{
		{po.EntityModel{1, time.Now(), time.Now()}, 1, 123, "真棒", '0'},
		{po.EntityModel{2, time.Now(), time.Now()}, 2, 456, "垃圾", '0'},
		{po.EntityModel{6, time.Now(), time.Now()}, 1, 666, "别说话", '0'},
		{po.EntityModel{7, time.Now(), time.Now()}, 1, 777, "尼玛", '0'},
		{po.EntityModel{8, time.Now(), time.Now()}, 1, 888, "发", '0'},
		{po.EntityModel{9, time.Now(), time.Now()}, 1, 999, "还是", '0'},
		{po.EntityModel{10, time.Now(), time.Now()}, 1, 000, "含税", '0'},
	}
	var co []bo.Comment = make([]bo.Comment, len(coms))
	GetCommentBOS(&coms, &co)
	fmt.Println("容量足够的情况下：")
	for _, c := range co {
		fmt.Println(c)
	}
	var co1 []bo.Comment
	GetCommentBOS(&coms, &co1)
	fmt.Print("没有初始化的情况下：")
	fmt.Println(len(co1))
	for _, c := range co1 {
		fmt.Println(c)
	}
	var co2 []bo.Comment = make([]bo.Comment, 0)
	GetCommentBOS(&coms, &co2)
	fmt.Print("容量不够的情况下：")
	fmt.Println(len(co2))
	for _, c := range co2 {
		fmt.Println(c)
	}
}
func TestGetUserBOS(t *testing.T) {
	var users []po.User = []po.User{
		{po.EntityModel{1, time.Now(), time.Now()}, "张三", "siw", 1, 2},
		{po.EntityModel{2, time.Now(), time.Now()}, "李斯", "siw", 3, 4},
		{po.EntityModel{3, time.Now(), time.Now()}, "张静怡", "siw", 5, 6},
		{po.EntityModel{4, time.Now(), time.Now()}, "张诺", "siw", 7, 8},
		{po.EntityModel{77, time.Now(), time.Now()}, "李文静", "siw", 99, 90},
	}
	var dest []bo.User = make([]bo.User, 5)
	GetUserBOS(&users, &dest)
	fmt.Print("容量充足的情况下：")
	fmt.Println(len(dest))
	for _, des := range dest {
		fmt.Println(des)
	}

	var dest1 []bo.User
	GetUserBOS(&users, &dest1)
	fmt.Print("没初始化的情况下：")
	fmt.Println(len(dest1))
	for _, des := range dest1 {
		fmt.Println(des)
	}

	var dest2 []bo.User = make([]bo.User, 0)
	GetUserBOS(&users, &dest2)
	fmt.Print("容量不足的情况下：")
	fmt.Println(len(dest2))
	for _, des := range dest2 {
		fmt.Println(des)
	}
}
func TestGetUserBO(t *testing.T) {
	src := po.User{
		po.EntityModel{1, time.Now(), time.Now()}, "李文静", "siw", 99, 90}
	var dest bo.User
	GetUserBO(&src, &dest)
	fmt.Println(dest)
}
func TestGetVideoBOS(t *testing.T) {
	var videos []po.Video = []po.Video{
		{po.EntityModel{1, time.Now(), time.Now()}, "xxx.com", "sddl.cn", 666, 0, 1, "xx"},
		{po.EntityModel{2, time.Now(), time.Now()}, "x1x.com", "s5l.cn", 0, 0, 2, "xx"},
		{po.EntityModel{3, time.Now(), time.Now()}, "x2x.com", "sd34l.cn", 0, 0, 6, "xx"},
		{po.EntityModel{5, time.Now(), time.Now()}, "x3x.com", "s45l.cn", 666, 0, 2, "xx"},
		{po.EntityModel{7, time.Now(), time.Now()}, "x4x.com", "s7dl.cn", 0, 0, 1, "xx"},
		{po.EntityModel{8, time.Now(), time.Now()}, "x5x.com", "sd56l.cn", 0, 0, 4, "xx"},
	}
	var dest []bo.Video = make([]bo.Video, len(videos))
	GetVideoBOS(&videos, &dest)
	fmt.Print("容量充足情况下：")
	fmt.Println(len(dest))
	for _, des := range dest {
		fmt.Println(des)
	}
	var dest1 []bo.Video
	GetVideoBOS(&videos, &dest1)
	fmt.Print("未初始化情况下：")
	fmt.Println(len(dest1))
	for _, des := range dest1 {
		fmt.Println(des)
	}
	var dest2 []bo.Video = make([]bo.Video, 0)
	GetVideoBOS(&videos, &dest2)
	fmt.Print("容量不足情况下：")
	fmt.Println(len(dest2))
	for _, des := range dest2 {
		fmt.Println(des)
	}
}
