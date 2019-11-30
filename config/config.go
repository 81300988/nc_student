package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Config Configchema

func init() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err1 := viper.Unmarshal(&Config)
	if err1 != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

type Configchema struct {
	Mongo struct {
		Uri      string `mapstructure:"Uri"`
		Host     string `mapstructure:"Host"`
		Usename  string `mapstructure:"Usename"`
		Password string `mapstructure:"Password"`
	} `mapstructure:"MongoDB"`
	JWTSecret struct {
		JWTKey string
	}
}
