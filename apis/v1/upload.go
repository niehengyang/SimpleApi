package v1

import (
	Middlewares "SimpleApi/middleware"
	Models "SimpleApi/models"
	"github.com/gin-gonic/gin"
)

// @Summary 上传文件
// @Description 上传文件接口
// @Tags 文件传输接口
// @Accept multipart/form-data
// @Produce multipart/form-data
// @Param token header string false "用户令牌"
// @Param file formData file true "文件"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/uploadfile [post]
func UpFile(c *gin.Context) {

	uploadResult, err := Models.UpLoadFile(c)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, err)
		return
	}

	//保存地址
	fpath := uploadResult.Path
	pathDir := fpath[7 : len(fpath)-0]
	uploadResult.Path = pathDir

	Middlewares.ResponseSuccess(c, uploadResult, "上传成功")
	return
}
