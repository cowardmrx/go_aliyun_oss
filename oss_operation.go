package go_aliyun_oss

import (
	"bytes"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliOssClient struct {
	Domain           string
	OriginalFileName bool
	Client           *oss.Bucket
}

type ossResponse struct {
	Host      string
	LongPath  string
	ShortPath string
	FileName  string
}

// Put 推送文件到oss
// params:  ossDir string  `oss dir [要推送到的oss目录]`  example: test/20201121/
// params:  file interface `upload file resource [文件资源]`
// return string  `oss file accessible uri [可访问地址]`
func (client *AliOssClient) Put(ossDir string, file interface{}, fileType string) (*ossResponse, error) {
	//file to []byte
	//文件转字节流
	uploadFile := &OssFile{
		File:     file,
		FileType: fileType,
	}

	ossFile, err := uploadFile.FileTypeTransForm()

	if err != nil {
		return nil, err
	}

	// 最终的oss名称
	var ossFileName string

	//ossPath = oss dir + upload file name
	//example: oss dir is diy ==== test/20201121/
	//time.Now().Format("20060102")
	//ossPath := path + fileName
	var ossPath string

	//judge is use origin file name if false fileName = fileNewName (is a only name) else file init name
	if client.OriginalFileName == false {
		ossPath = ossDir + ossFile.FileNewName
		ossFileName = ossFile.FileNewName
	} else {
		ossPath = ossDir + ossFile.FileOldName
		ossFileName = ossFile.FileOldName
	}

	//upload file to oss
	err = client.Client.PutObject(ossPath, bytes.NewReader(ossFile.FileByte))

	if err != nil {
		return nil, errors.New("put file to oss failed:" + err.Error())
	}

	return &ossResponse{
		Host:      client.Domain,
		LongPath:  client.Domain + "/" + ossPath,
		ShortPath: ossPath,
		FileName:  ossFileName,
	}, nil
}

// HasExists 校验文件是否已经存在
// check file already exists in oss server
// params: ossFilePath	string 	`file oss path [文件的oss的路径]`
func (client *AliOssClient) HasExists(ossFilePath string) (bool, error) {

	//oss check fun
	isExists, err := client.Client.IsObjectExist(ossFilePath)

	if err != nil {
		return false, errors.New("check file in oss is exists failed:" + err.Error())
	}

	return isExists, nil
}

// Delete 删除文件-单文件删除
// delete one file in oss
// params ossPath string `file oss path [文件的oss路径]`
// return bool
func (client *AliOssClient) Delete(ossFilePath string) (bool, error) {

	//oss delete one file fun
	err := client.Client.DeleteObject(ossFilePath)

	if err != nil {
		return false, errors.New("delete file " + ossFilePath + " failed:" + err.Error())
	}

	return true, nil
}

// DeleteMore 删除文件-多文件删除
// delete more file in oss
// params ossPath []string `file oss path array [文件的oss路径数组]`
// return bool
func (client *AliOssClient) DeleteMore(ossFilePath []string) (bool, error) {
	//oss delete more file fun
	_, err := client.Client.DeleteObjects(ossFilePath)

	if err != nil {
		return false, errors.New("delete more file in oss failed:" + err.Error())
	}

	return true, nil
}

// GetTemporaryUrl 获取文件临时地址
// path string 文件路径
// expireInSecond int64 多久后过期 单位: 秒，默认 60
func (client *AliOssClient) GetTemporaryUrl(path string, expireInSecond int64) (string, error) {

	var expireTime int64

	if expireInSecond <= 0 {
		expireTime = 60
	} else {
		expireTime = expireInSecond
	}

	signUrl, err := client.Client.SignURL(path, oss.HTTPGet, expireTime)

	if err != nil {
		return "", errors.New("generate sign url failed:" + err.Error())
	}

	return signUrl, nil

}
