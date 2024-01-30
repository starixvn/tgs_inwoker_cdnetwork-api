package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"openApi-authentication/common/constant"
	"openApi-authentication/common/model"
	"openApi-authentication/common/util"
	"time"
)

type BasicConfig struct {
	UserName string
	ApiKey   string
	Uri      string
	Method   string
	EndPoint string
}

func TransferRequestMsg(config BasicConfig) model.HttpRequestMsg {
	var requestMsg = model.HttpRequestMsg{Params: map[string]string{}, Headers: map[string]string{}}
	requestMsg.Uri = config.Uri
	requestMsg.Method = config.Method
	requestMsg.Url = constant.HttpRequestPrefix + config.Uri
	if len(config.EndPoint) == 0 || "{endPoint}" == config.EndPoint {
		requestMsg.Host = constant.HttpRequestDomain
		requestMsg.Url = "https://" + constant.HttpRequestDomain + config.Uri
	} else {
		requestMsg.Host = config.EndPoint
		requestMsg.Url = "https://" + config.EndPoint + config.Uri
	}
	return requestMsg
}

func Invoke(config BasicConfig, jsonStr string) string {
	var requestMsg = TransferRequestMsg(config)

	//jsonStu, err := json.Marshal(object)
	//if err != nil {
	//	fmt.Println("生成JSON字符串错误")
	//}

	if config.Method == "POST" || config.Method == "PUT" || config.Method == "PATCH" || config.Method == "DELETE" {
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
