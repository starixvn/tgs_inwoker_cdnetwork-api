package util

import (
	"cdnetwork/pkg/log"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type CdnConfig struct {
	UserName                  string `yaml:"user_name"`
	ApiKey                    string `yaml:"api_key"`
	EndPoint                  string `yaml:"end_point"`
	ChromeDriverPath          string `yaml:"chrome_driver_path"`
	CdnLoginUrl               string `yaml:"cdn_login_url"`
	CdnCertificateCreationUrl string `yaml:"cdn_certificate_creation_url"`
	CdnUsername               string `yaml:"cdn_username"`
	CdnPassword               string `yaml:"cdn_password"`
	TelegramBotToken          string `yaml:"telegram_bot_token"`
	TelegramChatId            string `yaml:"telegram_chat_id"`
	CloudflareToken           string `yaml:"cloudflare_token"`
	TelegramWebhook           string `yaml:"telegram_webhook"`
	DnsContent                string `yaml:"dns_content"`
}

func GetConfig() CdnConfig {
	data, err := os.ReadFile(os.Getenv("CONFIGPATH"))

	if err != nil {
		log.LogFatal(fmt.Sprintf("Fail to load config: %s  %v", os.Getenv("CONFIGPATH"), err))
		panic(err)
	}

	// create a person struct and deserialize the data into that struct
	var config CdnConfig

	if err := yaml.Unmarshal(data, &config); err != nil {
		log.LogFatal(fmt.Sprintf("Fail to parse config: %v", err))
		panic(err)
	}
	return config
}
