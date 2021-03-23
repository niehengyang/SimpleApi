package main

import (
	mysql "SimpleApi/databses"
	"SimpleApi/databses/seeder"
	Models "SimpleApi/models"
	"SimpleApi/pkg/setting"
	"SimpleApi/routers"
	"fmt"
	"net/http"
)

// @title 自定义项目API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email 790227542@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8088
// @BasePath

func main() {

	db := mysql.InitDB()
	Models.StartMigrate(db)
	seeder.SeedUser(db)
	seeder.SeedPermissions(db)
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), //监听的TCP地址，格式为:8000
		Handler:        router,                               //http句柄，实质为ServeHTTP,用于处理程序响应HTTP请求
		ReadTimeout:    setting.ReadTimeOut,                  //读取整个文件的最大持续时间s
		WriteTimeout:   setting.WriteTimeOut,                 //允许写入的最大时间s
		MaxHeaderBytes: 1 << 20,                              //请求主体的最大字节数
	}
	s.ListenAndServe()

	defer db.Close()
}
