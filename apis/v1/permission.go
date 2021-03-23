package v1

import (
	Middlewares "SimpleApi/middleware"
	Models "SimpleApi/models"
	jwt "SimpleApi/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// @Summary 查询权限列表
// @Description 查询权限列表接口
// @Tags 权限相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param page query int false "页码"
// @Param pageSize query int false "单页数量"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getpermissionlist [get]
func GetPermissionList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	users := Models.GetPermissionList(jwt.StringToInt(page), jwt.StringToInt(pageSize))

	Middlewares.ResponseSuccess(c, users, "查询成功")
}

// @Summary 查询所有权限
// @Description 查询所有权限接口
// @Tags 权限相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getpermissionall [get]
func GetAllPermissions(c *gin.Context) {

	permissions, err := Models.GetAllPermissions()
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("查询失败"))
		return
	}

	Middlewares.ResponseSuccess(c, permissions, "查询成功")
}
