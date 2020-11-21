package main

import "fmt"

var ossClient *AliOssClient

func init()  {
	ossConfig := &AliOssConfig{
		EndPoint: "oss-cn-shenzhen.aliyuncs.com",
		AccessKeyId: "",
		AccessKeySecret: "",
		BucketName: "",
	}

	ossClient = ossConfig.CreateOssConnect()
}

func main() {

	uri := ossClient.Put("logo/","./File/3HaqWaOzJWD86DDvZD9Pmn9VUEOBOBbuWackGOXb (2).jpeg")

	fmt.Println(uri)
}
