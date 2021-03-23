package apis

import (
	Middlewares "SimpleApi/middleware"
	Models "SimpleApi/models"
	jwt "SimpleApi/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录
// @Description 用户登录接口
// @Tags 鉴权相关接口
// @Accept application/json
// @Produce application/json
// @Param param body Models.LoginInput true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var loginInput Models.LoginInput
	err := c.ShouldBindJSON(&loginInput)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("数据获取失败"))
		return
	}

	//参数校验
	//TODO: 参数检验

	token, err := Models.Login(loginInput.Username, loginInput.Password)
	if err != nil {
		if err.Error() == "record not found" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("该用户不存在"))
			return
		} else if err.Error() == "invalid password" {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("密码校验不通过"))
			return
		} else {
			Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("登录错误"))
			return
		}
	}
	Middlewares.ResponseSuccess(c, token, "登录成功")
	return
}

// @Summary 注销登录
// @Description 注销登录接口
// @Tags 鉴权相关接口
// @Accept application/json
// @Produce application/json
// @Param token header string false "用户令牌"
// @Security ApiKeyAuth
// @Success 200 {object} Middlewares.Response
// @Router /auth/logout [get]
func Logout(c *gin.Context) {
	Middlewares.ResponseSuccess(c, nil, "注销成功")
}

//更新token
func InitiativeExpire(c *gin.Context) {
	j := jwt.NewJWT()
	newToken, err := j.RefreshToken(Middlewares.AuthToken)
	if err != nil {
		Middlewares.ResponseError(c, Middlewares.RequestFail, errors.New("token更新失败"))
		return
	}
	Middlewares.ResponseSuccess(c, newToken, "更新成功")
}
