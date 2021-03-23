package Models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"path"
	"time"
)

func UpLoadFile(c *gin.Context) (uploadResult FileForm, err error) {
	var result FileForm

	h, _ := c.FormFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".zip":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		return result, errors.New("后缀名不符合上传要求")
	}

	//创建目录地址
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")

	//创建目录
	err1 := os.MkdirAll(uploadDir, 0777)
	if err1 != nil {
		return result, errors.New("无权限创建目录")
	}

	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))

	fileName := fmt.Sprintf("%x", hashName) + ext

	fpath := uploadDir + fileName

	err2 := c.SaveUploadedFile(h, fpath)
	if err2 != nil {
		return result, errors.New("保存的时候报了错")
	}

	result.Path = fpath
	result.Name = fileName
	result.OriginName = h.Filename
	result.HashName = fmt.Sprintf("%x", hashName)

	return result, nil
}
