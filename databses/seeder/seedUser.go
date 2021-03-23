package seeder

import (
	Models "SimpleApi/models"
	"github.com/jinzhu/gorm"
)

func SeedUser(db *gorm.DB) {
	//填充数据
	result := db.Find(&Models.User{}, "1")
	if result.Error != nil {
		var rootInfo Models.UserForm
		var roles []int
		rootInfo.User.Name = "超级管理员"
		rootInfo.User.Username = "root"
		rootInfo.User.Password = "123456"
		rootInfo.User.Email = "root@threejs.com"
		rootInfo.User.Type = "00"
		rootInfo.User.Phone = "13999999999"
		rootInfo.User.Describe = "就是超管"
		rootInfo.Roles = roles

		Models.CreateUser(rootInfo)
	}
}
