package util

import (
	"cdnetwork/pkg/log"
	"fmt"

	"github.com/spf13/viper"
)

type CdnConfig struct {
	UserName                  string
	ApiKey                    string
	EndPoint                  string
	ChromeDriverPath          string
	CdnLoginUrl               string
	CdnCertificateCreationUrl string
	CdnUsername               string
	CdnPassword               string
}

func GetConfig() CdnConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		log.LogFatal(fmt.Sprintf("Error reading config file:", err))
	}

	var config CdnConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		log.LogFatal(fmt.Sprintf("Error unmarshalling config:", err))
	}
	return config
}
