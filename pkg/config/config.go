package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

var config *viper.Viper

func LoadConfig(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("./app")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
    fmt.Println("config", config)
	return config
}
