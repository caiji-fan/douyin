/**
 * @Author yg
 * @Date 2022-05-20
 * @Description
 **/
package config

type jwt struct {
	ExpireTime int    `yaml:"expire-time"`
	SecretKey  string `yaml:"secret-key"`
}
