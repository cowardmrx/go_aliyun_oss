package go_aliyun_oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliOssClient struct {
	Domain string
	OriginalFileName bool
	Client *oss.Bucket
}

type ossResponse struct {
	Host string
	LongPath string
	ShortPath string
	FileName string
}


//	@method: Put
//	@description: 推送文件到oss
//	@author: mr.x 2021-06-19 00:29:08
//	@param: ossDir string	`[要推送到的oss目录]`  example: test/20201121/
//	@param: file interface{}	`upload file resource [文件资源]`
//	@param: fileType string	`文件类型`
//	@return: *ossResponse 返回oss可访问地址等
func (client *AliOssClient) Put(ossDir string, file interface{},fileType string) *ossResponse {
	//file to []byte
	//文件转字节流
	uploadFile := &OssFile{
		File: file,
		FileType: fileType,
	}

	ossFile,err := uploadFile.FileTypeTransForm()

	if err != nil {
		panic("transfer file failed" + err.Error())
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
	err = client.Client.PutObject(ossPath,bytes.NewReader(ossFile.FileByte))

	if err != nil {
		panic("put file to oss failed:" + err.Error())
	}

	return &ossResponse{
		Host: client.Domain,
		LongPath: client.Domain + "/" + ossPath,
		ShortPath: ossPath,
		FileName: ossFileName,
	}
}

//	@method: HasExists
//	@description: 校验文件是否已经存在
//	@author: mr.x 2021-06-19 00:30:21
//	@param: ossFilePath string	file oss path [文件的oss的路径]
//	@return: bool
func (client *AliOssClient) HasExists(ossFilePath string) bool {

	//oss check fun
	isExists,err := client.Client.IsObjectExist(ossFilePath)

	if err != nil {
		panic("check file in oss is exists failed:" + err.Error())
	}

	return isExists
}

//	@method: Delete
//	@description: 删除文件-单文件删除
//	@author: mr.x 2021-06-19 00:30:40
//	@param: ossFilePath string oss 可访问路径
//	@return: bool true - 删除成功 | false - 删除失败
func (client *AliOssClient) Delete(ossFilePath string) bool {

	//oss delete one file fun
	err := client.Client.DeleteObject(ossFilePath)

	if err != nil {
		panic("delete file "+ ossFilePath +" failed:" + err.Error())
	}

	return true
}

//	@method: DeleteMore
//	@description: 删除文件-多文件删除
//	@author: mr.x 2021-06-19 00:30:56
//	@param: ossFilePath []string
//	@return: bool true - 批量删除成功 | false - 批量删除失败
func (client *AliOssClient) DeleteMore(ossFilePath []string) bool {
	//oss delete more file fun
	_,err := client.Client.DeleteObjects(ossFilePath)

	if err != nil {
		panic("delete more file in oss failed:" + err.Error())
	}

	return true
}


//	@method: GetTemporaryUrl
//	@description: 获取文件临时地址
//	@author: mr.x 2021-06-19 00:31:34
//	@param: path string 文件路径【段路径】
//	@param: expireInSecond int64 有效时间 秒 默认 60S
//	@return: string
func (client *AliOssClient) GetTemporaryUrl(path string,expireInSecond int64) string {

	var expireTime int64

	if expireInSecond <= 0 {
		expireTime = 60
	} else {
		expireTime = expireInSecond
	}

	signUrl,err := client.Client.SignURL(path,oss.HTTPGet,expireTime)

	if err != nil {
		panic("generate sign url failed:" + err.Error())
	}

	return signUrl

}