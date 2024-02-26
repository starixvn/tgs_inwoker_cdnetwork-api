package util

import (
	"bytes"
	"cdnetwork/pkg/log"
	"fmt"
	"os"

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
	TelegramBotToken          string
	TelegramChatId            string
}

func GetConfig() CdnConfig {

	os.Setenv("CONFIGPATH", "../../config/config.yaml")
	configPath := os.Getenv("CONFIGPATH")

	fmt.Sprintf("configPath: %s", configPath)

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.LogFatal(fmt.Sprintf("Error reading config file: %v", err))
	}

	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(fileContent))

	var config CdnConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		log.LogFatal(fmt.Sprintf("Error unmarshalling config:", err))
	}

	fmt.Println("UserName:", config.UserName)

	return config
}
