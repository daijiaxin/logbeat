package main

import (
	"log"
	"logbeat/internal/config"
	"logbeat/internal/filter"
	"logbeat/internal/input"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var options = config.Option{}

var configs = config.Config{}

func init() {
	_, err := flags.Parse(&options)
	if err != nil {
		os.Exit(1)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		DisableColors: true,
		ForceQuote:    true,
	})
}

func main() {
	viper.SetConfigFile(options.Config)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("load config failed: %v", err)
	}
	err = viper.Unmarshal(&configs)
	if err != nil {
		log.Fatalf("parser config failed: %v", err)
	}
	input.StartAll(configs.Inputs)

	for i := 0; i < 4; i++ {
		go filter.Start(i)
	}

	select {}
}
