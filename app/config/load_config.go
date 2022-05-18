package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var path = "\\resources\\application.yml"

var vip *viper.Viper

//项目启动时自动执行
func init() {
	rootPath, err := os.Getwd()
	ymlPath := rootPath + path
	vip = viper.New()
	vip.SetConfigFile(ymlPath)
	vip.SetConfigType("yaml")
	vip.AddConfigPath("config")
	err = vip.ReadInConfig()
	if err != nil {
		fmt.Println("err:vip.ReadInConfig()===>\n", err)
	}

}

//获取application.yml key对应 值
func Get(key string) string {
	return vip.GetString(key)
}
