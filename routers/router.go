package routers

import (
	"SimpleApi/apis"
	v1 "SimpleApi/apis/v1"
	_ "SimpleApi/docs"
	Middlewares "SimpleApi/middleware"
	"SimpleApi/pkg/setting"
	jwt "SimpleApi/pkg/utils"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	//swagger api文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	r.Use(Middlewares.Cors())

	// 日志记录到文件。
	r.Use(Middlewares.LoggerToFile())

	//登录注册
	auth := r.Group("/auth")
	{
		auth.POST("/login", apis.Login)
		auth.GET("/logout", apis.Logout, jwt.JWTAuth())
		//auth.POST("/register", Controllers.Register)
	}

	//v1版本api
	apiv1 := r.Group("/v1", jwt.JWTAuth())
	{

		//用户
		apiv1.POST("/createuser", v1.CreateUser)
		apiv1.PUT("/edituser/:id", v1.EditUser)
		apiv1.GET("/getuserlist", v1.GetUserList)
		apiv1.GET("/getuseritem/:id", v1.GetUserItem)
		apiv1.GET("/getloginuser", v1.GetUserInfo)
		apiv1.PUT("/bindroles/:id", v1.BindRoles)
		apiv1.DELETE("/deleteuser/:id", v1.DeleteUser)

		//权限
		apiv1.GET("/getpermissionlist", v1.GetPermissionList)
		apiv1.GET("/getpermissionall", v1.GetAllPermissions)

		//角色
		apiv1.POST("/createrole", v1.CreateRole)
		apiv1.PUT("/editrole/:id", v1.EditRole)
		apiv1.GET("/getroleitem/:id", v1.GetRoleItem)
		apiv1.GET("/getrolelist", v1.GetRoleList)
		apiv1.DELETE("/deleterole/:id", v1.DeleteRole)
		apiv1.PUT("/bindpermissions/:id", v1.BindPermissions)
		apiv1.GET("/getrolepermissions/:id", v1.GetRolePermissions)

		//上传文件
		apiv1.POST("/uploadfile", v1.UpFile)

		//cehp api
		apiv1.POST("/upfilebyceph", v1.UploadFileToCeph)
	}

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
