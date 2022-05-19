// Package encryptionutil
// @Author shaofan
// @Date 2022/5/14
// @DESC 加密工具
package encryptionutil

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

// Encryption 			加密
// src					需要加密的字符串
// @return 				加密后的字符串
func Encryption(src string) (string, error) {
	if src == "" {
		return "", errors.New("错误：密码为空")
	}
	data := []byte(src) //首先转换成字符串
	md5data := md5.Sum(data)
	return fmt.Sprintf("%x", md5data), nil
}

// EncryptionCompare 	加密对比
// src					原字符串
// encryptionString 	待对比的加密字符串
func EncryptionCompare(src string, encryptionString string) (bool, error) {
	if src == "" {
		return false, errors.New("错误：密码为空")
	}
	md5src, err := Encryption(src)
	if err != nil {
		return false, err
	}
	if strings.EqualFold(md5src, encryptionString) {
		return true, nil
	}
	return false, errors.New("用户名或密码输入错误，请重试")
}
