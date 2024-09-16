package config

import "github.com/spf13/viper"

func New() Config {
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config
}
