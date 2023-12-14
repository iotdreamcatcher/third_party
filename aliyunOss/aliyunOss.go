package aliyunOss

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"time"
)

// 请填写您的AccessKeyId。
//var accessKeyId string = "<yourAccessKeyId>"
//// 请填写您的AccessKeySecret。
//var accessKeySecret string = "<yourAccessKeySecret>"
//// host的格式为 bucketname.endpoint ，请替换为您的真实信息。
//var host string = "http://bucket-name.oss-cn-hangzhou.aliyuncs.com"
//// callbackUrl为 上传回调服务器的URL，请将下面的IP和Port配置为您自己的真实信息。
//var callbackUrl string = "http://88.88.88.88:8888";
//// 用户上传文件时指定的前缀。
//var upload_dir string = "upload/"
//var expire_time int64 = 30

type AliyunConf struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Host            string `json:"host"`
	TenantId        uint   `json:"tenant_id"`
	IsCallback      bool   `json:"is_callback"`
	Prefix          string `json:"prefix"`
	UploadDir       string `json:"upload_dir"`
	Key             string `json:"key"`
	CallbackUrl     string `json:"callback_url"`
	ExpireTime      int64  `json:"expire_time"`
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

func (conf AliyunConf) GetPolicyToken() (*PolicyToken, error) {
	var policyToken PolicyToken
	now := time.Now().Unix()
	expire_end := now + conf.ExpireTime
	//var tokenExpire = get_gmt_iso8601(expire_end)

	uploadUrl := fmt.Sprintf("%s/TID%d/%s/%s", conf.Prefix, conf.TenantId, conf.UploadDir, conf.Key)

	//create post policy json
	var config ConfigStruct
	config.Expiration = get_gmt_iso8601(expire_end)
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, uploadUrl)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(conf.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	if conf.IsCallback {
		var callbackParam CallbackParam
		callbackParam.CallbackUrl = conf.CallbackUrl
		callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
		callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
		callback_str, err := json.Marshal(callbackParam)
		if err != nil {
			return nil, err
		}
		callbackBase64 := base64.StdEncoding.EncodeToString(callback_str)
		policyToken.Callback = string(callbackBase64)
	}

	policyToken.AccessKeyId = conf.AccessKeyId
	policyToken.Host = conf.Host
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = uploadUrl
	policyToken.Policy = string(debyte)
	//response, err := json.Marshal(policyToken)
	//if err != nil {
	//	return "", err
	//}
	return &policyToken, nil
}

func get_gmt_iso8601(expire_end int64) string {
	var tokenExpire = time.Unix(expire_end, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}
