package main

import (
	"fmt"
	"github.com/mohuishou/apollo-viper-provider"
	_ "github.com/mohuishou/apollo-viper-provider"
	"github.com/spf13/viper"

)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("apollo.server", "http://129.211.63.96:8080")
	viper.SetDefault("apollo.app_id", "SampleApp")
	v, _ := apollo.InitViper(viper.Sub("apollo"))

	// 获取配置信息
	fmt.Println(v.GetInt("timeout"))

}
