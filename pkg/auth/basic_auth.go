package auth

import (
	"cdnetwork/config"
	"cdnetwork/pkg/constant"
	"cdnetwork/pkg/model"
	"cdnetwork/pkg/util"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type BasicConfig struct {
	Uri    string
	Method string
}

func TransferRequestMsg(parameter BasicConfig, config config.CdnConfig) model.HttpRequestMsg {

	var requestMsg = model.HttpRequestMsg{Params: map[string]string{}, Headers: map[string]string{}}
	requestMsg.Uri = parameter.Uri
	requestMsg.Method = parameter.Method
	requestMsg.Url = constant.HttpRequestPrefix + parameter.Uri
	if len(config.EndPoint) == 0 || "{endPoint}" == config.EndPoint {
		requestMsg.Host = constant.HttpRequestDomain
		requestMsg.Url = "https://" + constant.HttpRequestDomain + parameter.Uri
	} else {
		requestMsg.Host = config.EndPoint
		requestMsg.Url = "https://" + config.EndPoint + parameter.Uri
	}
	return requestMsg
}

func Invoke(url BasicConfig, jsonStr string) string {
	config := loadConfig()

	var requestMsg = TransferRequestMsg(url, config)

	if url.Method == "POST" || url.Method == "PUT" || url.Method == "PATCH" || url.Method == "DELETE" {
		requestMsg.Body = jsonStr
	}

	dateStr := getRfc1123Date()
	authorization := "Basic " + authorization(config.UserName, hmacsha1(dateStr, config.ApiKey))
	requestMsg.Headers["Date"] = dateStr
	requestMsg.Headers["Host"] = requestMsg.Host
	requestMsg.Headers["Content-Type"] = constant.ApplicationJson
	requestMsg.Headers[constant.Authorization] = authorization
	requestMsg.Headers[constant.XCncAuthMethod] = constant.BASIC
	return util.Call(requestMsg)
}

func authorization(userName string, passwd string) string {
	result := base64.StdEncoding.EncodeToString([]byte(userName + ":" + passwd))
	return result
}

func hmacsha1(date string, apikey string) string {
	key := []byte(apikey)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(date))
	value := mac.Sum(nil)
	result := base64.StdEncoding.EncodeToString(value)
	return result
}

func getRfc1123Date() string {
	date := time.Now().UTC() // UTC time
	return date.Format(http.TimeFormat)
}

func loadConfig() config.CdnConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../config")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		panic(fmt.Sprintf("Error reading config file:", err))
	}

	var config config.CdnConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Error unmarshalling config:", err)
		panic(fmt.Sprintf("Error unmarshalling config:", err))
	}
	return config
}
