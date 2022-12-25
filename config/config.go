package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type featureFlag struct {
	ConfigFilePath string
	UserDefault    string
}

type Config struct {
	FeatureFlag featureFlag
}

func LoadConfig() (*Config, error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env") // name of Config file (without extension)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	c := &Config{
		FeatureFlag: featureFlag{
			ConfigFilePath: getRequiredString("CONFIG_FILE_PATH"),
			UserDefault:    getRequiredString("USER_DEFAULT"),
		},
	}

	return c, nil
}

func getRequiredString(key string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}

	panic(fmt.Errorf("KEY %s IS MISSING", key))
}
