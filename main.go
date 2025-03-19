package main

import (
	"GZCTFBOT/config"
	"GZCTFBOT/control"

	"github.com/spf13/viper"
)

func main() {
	// Init config
	config.InitConfig()

	control.StartEvent(viper.GetString("Global.BaseURL"), viper.GetString("Global.GameID"), viper.GetInt("Global.RequestInterval"))
}
