package go_aliyun_oss

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPut(t *testing.T)  {
	ossConfig := &AliOssConfig{
		EndPoint: "",
		AccessKeyId: "",
		AccessKeySecret: "",
		BucketName: "",
	}

	client := ossConfig.CreateOssConnect()

	uri := client.Put("logo/","./File/3HaqWaOzJWD86DDvZD9Pmn9VUEOBOBbuWackGOXb (2).jpeg")

	fmt.Println(uri)
}

func TestPutBase64(t *testing.T)  {
	ossConfig := &AliOssConfig{
		EndPoint: "oss-cn-shenzhen.aliyuncs.com",
		AccessKeyId: "LTAI4FwmLMvboJf97riakG9e",
		AccessKeySecret: "1dclKyqjiI1SDfr30IDvO0eZmcJLo3",
		BucketName: "test-myself-bucket",
	}

	client := ossConfig.CreateOssConnect()

	// 读取base
	file,_ := ioutil.ReadFile("./File/3HaqWaOzJWD86DDvZD9Pmn9VUEOBOBbuWackGOXb (2).jpeg")


	fileStr := base64.StdEncoding.EncodeToString(file)

	bat,_ := base64.StdEncoding.DecodeString(fileStr)


	uri := client.Put("logo/",bat)
	//
	fmt.Println(uri)
}

func TestExists(t *testing.T)  {
	ossConfig := &AliOssConfig{
		EndPoint: "oss-cn-shenzhen.aliyuncs.com",
		AccessKeyId: "",
		AccessKeySecret: "",
		BucketName: "",
	}

	client := ossConfig.CreateOssConnect()

	isExists := client.HasExists("logo/a82bbd10-bb3f-5744-8843-5ef0d06c3b23.jpeg")

	fmt.Println(isExists)
}

func TestDelete(t *testing.T) {
	ossConfig := &AliOssConfig{
		EndPoint: "oss-cn-shenzhen.aliyuncs.com",
		AccessKeyId: "",
		AccessKeySecret: "",
		BucketName: "",
	}

	client := ossConfig.CreateOssConnect()

	deleteRes := client.Delete("logo/41e6ddf4-fe9a-53c3-8994-0a69aba031c7.jpeg")

	fmt.Println(deleteRes)
}

func TestAliOssClient_DeleteMore(t *testing.T) {
	ossConfig := &AliOssConfig{
		EndPoint: "oss-cn-shenzhen.aliyuncs.com",
		AccessKeyId: "",
		AccessKeySecret: "",
		BucketName: "",
	}

	client := ossConfig.CreateOssConnect()

	deleteRes := client.DeleteMore([]string{
		"logo/b9f775db-8eb5-5652-86c9-5322ff4ba212.jpeg",
		"logo/a82bbd10-bb3f-5744-8843-5ef0d06c3b23.jpeg",
	})

	fmt.Println(deleteRes)
}

