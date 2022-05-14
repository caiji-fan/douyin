// Package jwtutil
// @Author shaofan
// @Date 2022/5/13
// @DESC jwt生成与校验工具
package jwtutil

// CreateJWT 	生成token
// id 			用户id
// @return 		token序列
func CreateJWT(id int) (string, error) {
	return "", nil
}

// ParseJWT 	解析token
// token 		需要解析的token
// @return 		如果token可用，则返回token中的用户id
func ParseJWT(token string) (int, error) {
	return 0, nil
}
