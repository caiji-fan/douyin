// Package config
// @Author shaofan
// @Date 2022/5/13
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// Config 配置信息
type config struct {
	DB     database `yaml:"database"`
	Redis  redis    `yaml:"redis"`
	Rabbit rabbit   `yaml:"rabbit"`
	Server server   `yaml:"server"`
}

// Config 全局配置实例
var Config *config

// 读取yml文件
func readConfig() {
	file, err := ioutil.ReadFile("./config/config.yml") //启动项目测试时用
	//file, err := ioutil.ReadFile("../../config/config.yml") //单元测试时用
	if err != nil {
		log.Fatalln("读取文件config.yml发生错误", err)
		return
	}
	if yaml.Unmarshal(file, Config) != nil {
		log.Fatalln("解析文件config.yml发生错误", err)
		return
	}
}

// Init 配置初始化
func Init() {
	Config = &config{}
	readConfig()
}
