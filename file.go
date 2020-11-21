package aliyun_oss

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

type OssFile struct {
	FileByte []byte
	FileOldName string
	FileNewName string
	FileType string
	File	interface{}
}

type OssFileInterface interface {
	FileTypeTransForm() (*OssFile,error)
	GetFileType() *OssFile
}

// file type transform
//@title 文件类型转换
func (ossFile *OssFile) FileTypeTransForm() (*OssFile,error) {
	var err error

	switch ossFile.File.(type) {
	case *os.File:
		ossFile.FileByte,err = ioutil.ReadAll(ossFile.File.(*os.File))

		if err != nil {
			panic("read os type file failed:" + err.Error())
		}

		_,ossFile.FileOldName = filepath.Split(ossFile.File.(*os.File).Name())

	case *multipart.FileHeader:

		fileResources,err := ossFile.File.(*multipart.FileHeader).Open()

		if err != nil {
			panic("open multipart file failed:" + err.Error())
		}

		defer fileResources.Close()

		ossFile.FileByte,err = ioutil.ReadAll(fileResources)

		if err != nil {
			panic("read multipart file failed:" + err.Error())
		}

		ossFile.FileOldName = ossFile.File.(*multipart.FileHeader).Filename

	case string:
		newFile,err := os.Open(ossFile.File.(string))

		if err != nil {
			panic("open file path failed:" + err.Error())
		}

		defer newFile.Close()

		ossFile.FileByte,err = ioutil.ReadAll(newFile)

		_,ossFile.FileOldName = filepath.Split(newFile.Name())

	default:
		fmt.Println(reflect.TypeOf(ossFile.File))
		panic("file type error" )
	}


	ossFile.GetFileType()

	return ossFile,nil
}

//split file type and generate file name
//截取文件类型
func (ossFile *OssFile) GetFileType() *OssFile {
	//from oldFileName split file type
	fileTypeSufIndex := strings.Index(ossFile.FileOldName,".")

	fileType := ossFile.FileOldName[fileTypeSufIndex:]

	ossFile.FileType = fileType

	//generate only file name
	ossFile.FileNewName = uuid.NewV5(uuid.NewV4(),ossFile.FileOldName).String() + ossFile.FileType

	return ossFile
}