package v1

import (
	Middlewares "SimpleApi/middleware"
	Models "SimpleApi/models"
	"SimpleApi/pkg/setting"
	"SimpleApi/servers/goCephApis"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func TraverseTheDocumentByCeph(c *gin.Context) {

}

func DownloadFileByCeph(c *gin.Context) {

}

func UploadFileToCeph(c *gin.Context) {

	uploadResult, err := Models.UpLoadFile(c)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, err)
		return
	}
	fpath := uploadResult.Path
	pathDir := fpath[7 : len(fpath)-0]

	//创建目录地址
	dstName := "/test/" + time.Now().Format("2006/01/02/") + uploadResult.OriginName

	cephStatus := goCephApis.Ceph.Init2(setting.CEPH_HOST, setting.CEPH_BUCKET_NAME, setting.CEPH_ACCESSKEY, setting.CEPH_SECRETKEY)
	if cephStatus != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("s3服务连接错误"))
		return
	}

	result := goCephApis.Ceph.Upload(pathDir, dstName)
	if result != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("文件上传失败"))
		return
	}

	Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("文件上传成功"))
	return
}

func DeleteFileForCeph(c *gin.Context) {

}

func DeleteAllFileForCeph(c *gin.Context) {

}
