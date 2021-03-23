package v1

import (
	Middlewares "SimpleApi/middleware"
	Models "SimpleApi/models"
	jwt "SimpleApi/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// @Summary 新建角色
// @Description 新建角色接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param data body Models.RoleForm true "角色结构体"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/createrole [post]
func CreateRole(c *gin.Context) {
	var data Models.Role
	err := c.ShouldBindJSON(&data)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	j := jwt.NewJWT()
	jwtClaims, _ := j.ParseToken(Middlewares.AuthToken)
	data.UserId = int(jwtClaims.Id)

	//创建
	modelId, result := Models.CreateRole(data)

	if result != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("创建失败"))
		return
	} else {
		Middlewares.ResponseSuccess(c, modelId, "创建成功")
	}
}

// @Summary 编辑角色
// @Description 编辑角色接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path int true "角色ID"
// @Param data body Models.RoleForm true "角色结构体"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/editrole [put]
func EditRole(c *gin.Context) {
	var data Models.Role

	roleId := c.Param("id")
	err := c.ShouldBindJSON(&data)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	//编辑
	result := Models.EditRole(jwt.StringToInt(roleId), data)

	if result {
		Middlewares.ResponseSuccess(c, roleId, "编辑成功")
	}
}

// @Summary 角色详情
// @Description 角色详情接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path integer true "角色ID"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getroleitem [get]
func GetRoleItem(c *gin.Context) {
	id := c.Param("id")

	//查找
	roleItem, status := Models.GetRoleItem(jwt.StringToInt(id))

	if status {
		Middlewares.ResponseSuccess(c, roleItem, "success")
		return
	} else {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("查询失败"))
		return
	}
}

// @Summary 查询角色列表
// @Description 查询角色列表接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param Name query string false "角色名称"
// @Param page query int false "页码"
// @Param pageSize query int false "单页数量"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getrolelist [get]
func GetRoleList(c *gin.Context) {

	name := c.DefaultQuery("name", "")
	page := c.DefaultQuery("page", "0")
	pageSize := c.DefaultQuery("pageSize", "10")

	models := Models.GetRoleList(name, jwt.StringToInt(page), jwt.StringToInt(pageSize))

	Middlewares.ResponseSuccess(c, models, "查询成功")
}

// @Summary 删除角色
// @Description 删除角色接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path integer true "角色ID"
// @Security ApiKeyAuths
// @Success 200 {object} Middlewares.Response
// @Router /v1/deleterole [delete]
func DeleteRole(c *gin.Context) {

	roleId := c.Param("id")
	err := Models.DeleteRole(roleId)

	if err != nil {
		if err.Error() == "role not find" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("该角色不存在"))
			return
		} else {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("删除失败"))
			return
		}

	}
	Middlewares.ResponseSuccess(c, nil, "删除成功")
	return
}

// @Summary 角色绑定权限
// @Description 角色绑定权限接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path int true "角色ID"
// @Param data body []int false "权限数组"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/bindpermissions [put]
func BindPermissions(c *gin.Context) {
	var permissionUids []string
	roleId := c.Param("id")

	err := c.ShouldBindJSON(&permissionUids)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	status := Models.BindPermissions(jwt.StringToInt(roleId), permissionUids)

	if status != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("绑定失败"))
	} else {
		Middlewares.ResponseSuccess(c, nil, "绑定成功")
	}
}

// @Summary 角色权限
// @Description 角色权限接口
// @Tags 角色相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Param id path integer true "角色ID"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /v1/getrolepermissions [get]
func GetRolePermissions(c *gin.Context) {
	id := c.Param("id")

	//查找
	permissions, err := Models.GetRolePermissions(jwt.StringToInt(id))

	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("查询失败"))
		return
	} else {
		Middlewares.ResponseSuccess(c, permissions, "success")
	}
}
