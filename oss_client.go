package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
)

type AliOssConfig struct {
	EndPoint 		string  `json:"end_point"`
	AccessKeyId 	string	`json:"access_id"`
	AccessKeySecret string	`json:"access_key"`
	BucketName 		string	`json:"bucket_name"`
	Domain 			string	`json:"domain"`
	OriginalFileName 	bool 	`json:"only_file_name"`
}

type AliOssConfigInterface interface {
	CheckConfig()
	CreateOssConnect() *AliOssClient
	GetAccessibleUrl() string
}

//check AliOssConfig value is exists
func (coon *AliOssConfig) CheckConfig() {
	//check endPoint
	if coon.EndPoint == "" || len(coon.EndPoint) <= 0 {
		panic("endPoint value can not empty")
	}

	//check endPoint http prefix if empty default http
	if strings.HasPrefix(coon.EndPoint,"https://") == false && strings.HasPrefix(coon.EndPoint,"http://") == false {
		coon.EndPoint = "http://" + coon.EndPoint
	}

	//check access secret
	if coon.AccessKeyId == "" || len(coon.AccessKeyId) <= 0 {
		panic("accessKeyId can not empty")
	}

	//check access key
	if coon.AccessKeySecret == "" || len(coon.AccessKeySecret) <= 0 {
		panic("accessKeySecret can not empty")
	}

	//check bucket name
	if coon.BucketName == "" || len(coon.BucketName) <= 0 {
		panic("bucketName can not empty")
	}

}

//en: create oss connect client
//创建阿里云oss 链接客户端
func (coon *AliOssConfig) CreateOssConnect() *AliOssClient {
	//config check
	coon.CheckConfig()

	//connect oss client/ init oss client
	//链接oss客户端/初始化oss客户端
	connect, err := oss.New(coon.EndPoint, coon.AccessKeyId, coon.AccessKeySecret)
	if err != nil {
		panic("connect oss client failed:" + err.Error())
	}

	//选择桶
	//choose oss bucket
	client,err := connect.Bucket(coon.BucketName)

	if err != nil {
		panic("choose bucket name failed:" + err.Error())
	}

	//拼接可访问地址
	//get accessible url
	var domain string
	domain = coon.GetAccessibleUrl()

	return &AliOssClient{
		Domain: domain,
		OriginalFileName: coon.OriginalFileName,
		Client: client,
	}
}

//get oss accessible url
//拼接阿里云oss可访问地址
func (coon *AliOssConfig) GetAccessibleUrl() string {
	var domain string

	//select endpoint http prefix
	//查找oss endpoint 的http 前缀
	endPointUriPrefixIndex := strings.Index(coon.EndPoint,"://")
	endPointUriPrefix := coon.EndPoint[:endPointUriPrefixIndex]

	//截取出endPoint
	//split oss public domain '://' length is 3
	endPoint := coon.EndPoint[endPointUriPrefixIndex + 3:]

	//judge accessible uri value is exists if not. accessible uri = endPointUriSuf + :// + bucketName + . + endPoint
	//example: http://test.oss-cn-shenzhen.aliyuncs.com
	//判断可访问地址是否配置
	if coon.Domain == "" || len(coon.Domain) <= 0 {
		domain = endPointUriPrefix + "://" + coon.BucketName + "." + endPoint

		//not exists
	} else {
		//judge domain is http prefix and https prefix
		if strings.HasPrefix("http://",coon.Domain) == false && strings.HasPrefix("https://",coon.Domain) == false {
			domain = endPointUriPrefix + "://" + coon.Domain
		} else {
			domain = coon.Domain
		}
	}

	return domain
}
