package main

import (
	"GZCTFBOT/action"
	"GZCTFBOT/config"
	"GZCTFBOT/control"

	"github.com/spf13/viper"
)

func main() {
	// Init config
	config.InitConfig()
	if viper.GetBool("Global.DingDIngBot") {
		action.NewDingDing()
	}
	if viper.GetBool("Global.QQBot") {
		action.NewQQ()
	}

	control.StartEvent(viper.GetString("Global.BaseURL"), viper.GetString("Global.GameID"), viper.GetInt("Global.RequestInterval"))

}
