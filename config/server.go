// Package config
// @Author shaofan
// @Date 2022/5/13
// @DESC
package config

type server struct {
	Port int    `yaml:"port"`
	Name string `yaml:"name"`
}
