package config

import (
	"github.com/spf13/viper"
)

func Load(envFile string) {
	viper.SetConfigFile(envFile)
	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetBool(key string) bool {
	return viper.GetBool(key)
}
