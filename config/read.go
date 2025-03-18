package config

import (
	"os"

	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() {
	workDir, _ := os.Getwd() //获取工作目录

	viper.SetConfigName("bot")   // 设置config名字
	viper.SetConfigType("yml")   //设置配置文件类型
	viper.AddConfigPath(workDir) // 设置配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

}
