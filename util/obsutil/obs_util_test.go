// Package obsutil
// @Author shaofan
// @Date 2022/5/16
package obsutil

import (
	"douyin/config"
	"fmt"
	"testing"
)

func TestMain(t *testing.M) {
	config.Init()
	t.Run()
}

func TestUpload(t *testing.T) {
	objectName, err := Upload("D:/douyin.sql", config.Config.Obs.Buckets.Video)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(objectName)
}

func TestGetName(t *testing.T) {
	//filePath := "d:/douyin.sql"
	//var objectName = strings.ToLower(filePath[strings.LastIndex(filePath, "."):])
	//var str = getUUID() + objectName
	//fmt.Println(strings.ReplaceAll(str, "-", ""))
	//
	//fmt.Println("https://" + config.Config.Obs.Buckets.Video + ".obs." + config.Config.Obs.Location)
	//objectName = strings.Replace(config.Config.Obs.EndPoint,
	//	"https://obs",
	//	"https://"+config.Config.Obs.Buckets.Video+".obs", 1)
	//fmt.Println(objectName)
}
