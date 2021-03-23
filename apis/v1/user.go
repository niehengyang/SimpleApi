package v1

import (
	Middlewares "SimpleApi/middleware"
	Models "SimpleApi/models"
	jwt "SimpleApi/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// @Summary 创建账号
// @Description 创建账号接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param data body Models.UserForm true "表单对象"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/createuser [post]
func CreateUser(c *gin.Context) {
	var userData Models.UserForm

	err := c.ShouldBindJSON(&userData)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}
	//创建
	userId, result := Models.CreateUser(userData)
	if result != nil {
		if result.Error() == "user search fail" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("用户查询失败"))
			return
		}
		if result.Error() == "user exit" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("用户名已存在"))
			return
		}
		if result.Error() == "user create fail" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("用户创建失败"))
			return
		}
	}
	Middlewares.ResponseSuccess(c, userId, "创建成功")
}

// @Summary 查询用户列表
// @Description 查询用户列表接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param username query string false "用户名"
// @Param page query int false "页码"
// @Param pageSize query int false "单页数量"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getuserlist [get]
func GetUserList(c *gin.Context) {

	username := c.DefaultQuery("username", "")
	page := c.DefaultQuery("page", "0")
	pageSize := c.DefaultQuery("pageSize", "10")

	users := Models.GetUserList(username, jwt.StringToInt(page), jwt.StringToInt(pageSize))

	Middlewares.ResponseSuccess(c, users, "查询成功")
}

// @Summary 查询用户详情
// @Description 查询用户详情接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户令牌"
// @Param id path integer true "用户ID"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getuseritem [get]
func GetUserItem(c *gin.Context) {
	id := c.Param("id")

	//查找
	userItem, status := Models.GetUserItem(jwt.StringToInt(id))

	if status != nil {
		if status.Error() == "user not find" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("该用户不存在"))
			return
		} else {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("查询失败"))
			return
		}
	}

	Middlewares.ResponseSuccess(c, userItem, "success")
	return
}

// @Summary 查询登录用户信息
// @Description 查询登录用户信息接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string true "用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getloginuser [get]
func GetUserInfo(c *gin.Context) {

	userItem, status := Models.GetLoginUser()

	if status {
		Middlewares.ResponseSuccess(c, userItem, "success")
		return
	} else {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("查询失败"))
		return
	}
}

// @Summary 用户绑定角色
// @Description 用户绑定角色接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path int true "用户ID"
// @Param data body []Models.RoleForm false "角色数组"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/bindroles [put]
func BindRoles(c *gin.Context) {
	var roles []Models.Role
	userId := c.Param("id")

	err := c.ShouldBindJSON(&roles)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	status := Models.BindRoles(jwt.StringToInt(userId), roles)

	if status != nil {
		if status.Error() == "user not find" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("该用户不存在"))
			return
		} else {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("绑定失败"))
			return
		}
	}
	Middlewares.ResponseSuccess(c, nil, "绑定成功")
}

// @Summary 编辑用户
// @Description 编辑用户接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path int true "用户ID"
// @Param data body Models.UserForm true "用户表单数据"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/edituser [put]
func EditUser(c *gin.Context) {
	var data Models.UserForm
	userId := c.Param("id")
	err := c.ShouldBindJSON(&data)

	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	//编辑
	result := Models.EditUser(jwt.StringToInt(userId), data)

	if result {
		Middlewares.ResponseSuccess(c, userId, "编辑成功")
		return
	}
}

// @Summary 删除用户
// @Description 删除用户接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path integer true "用户ID"
// @Security ApiKeyAuths
// @Success 200 {object} Middlewares.Response
// @Router /v1/deleteuser [delete]
func DeleteUser(c *gin.Context) {

	userId := c.Param("id")
	if userId == "" {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	err := Models.DeleteUser(jwt.StringToInt(userId))
	if err != nil {
		if err.Error() == "user not find" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("该用户不存在"))
			return
		} else {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("登录错误"))
			return
		}
	}
	Middlewares.ResponseSuccess(c, nil, "删除成功")
}
