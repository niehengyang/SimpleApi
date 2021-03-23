package seeder

import (
	Models "SimpleApi/models"
	"github.com/jinzhu/gorm"
)

func SeedPermissions(db *gorm.DB) {

	var perArr [8]Models.PermissionForm

	perArr[0].Permission = Models.Permission{Name: "首页", Parent: "0", Status: "1", Type: "01", Uid: "home", Url: "test", Icon: "test", Describe: "首页"}
	perArr[1].Permission = Models.Permission{Name: "系统管理", Parent: "0", Status: "1", Type: "01", Uid: "system", Url: "test", Icon: "test", Describe: "系统管理"}
	perArr[2].Permission = Models.Permission{Name: "用户管理", Parent: "2", Status: "1", Type: "02", Uid: "user", Url: "test", Icon: "test", Describe: "用户管理"}
	perArr[3].Permission = Models.Permission{Name: "日志管理", Parent: "2", Status: "1", Type: "02", Uid: "log", Url: "test", Icon: "test", Describe: "日志管理"}
	perArr[4].Permission = Models.Permission{Name: "模型管理", Parent: "0", Status: "1", Type: "01", Uid: "model", Url: "test", Icon: "test", Describe: "模型管理"}
	perArr[5].Permission = Models.Permission{Name: "应用管理", Parent: "0", Status: "1", Type: "01", Uid: "app", Url: "test", Icon: "test", Describe: "应用管理"}
	perArr[6].Permission = Models.Permission{Name: "角色管理", Parent: "2", Status: "1", Type: "02", Uid: "role", Url: "test", Icon: "test", Describe: "角色管理"}

	perArr[0].Permission.BaseModel.ID = 1
	perArr[1].Permission.BaseModel.ID = 2
	perArr[2].Permission.BaseModel.ID = 3
	perArr[3].Permission.BaseModel.ID = 4
	perArr[4].Permission.BaseModel.ID = 5
	perArr[5].Permission.BaseModel.ID = 6
	perArr[6].Permission.BaseModel.ID = 7

	for i := 0; i < len(perArr); i++ {
		db.FirstOrCreate(&perArr[i].Permission)
	}
}
