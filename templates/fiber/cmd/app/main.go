package main

import (
	"flag"
	"fmt"
	"path"

	"templates/internal/app"
	"templates/internal/domain"
	"templates/pkg/config/viper"
)

func main() {
	appConf := flag.String("config", "dev", "[prod,dev,locale]")
	flag.Parse()

	var appConfig domain.AppConfigs
	configFile := path.Join("etc", *appConf+".yaml")

	if err := viper.Parse(configFile, &appConfig); err != nil {
		panic(fmt.Sprintf("error parse configs error:%+v", err))
	}

	if err := app.Run(appConfig); err != nil {
		fmt.Printf("can't run app %+v", err)
	}
}
