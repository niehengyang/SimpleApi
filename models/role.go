package Models

import (
	Mysql "SimpleApi/databses"
	"errors"
)

func CreateRole(newRole Role) (uint, error) {
	result := Mysql.DB.Create(&newRole)
	id := newRole.ID
	if result.Error != nil {
		return 0, errors.New("create error")
	}
	return id, nil
}

func EditRole(roleId int, newRole Role) bool {

	var role Role
	Mysql.DB.Find(&role, roleId)

	Mysql.DB.Model(&role).Update(newRole)

	return true
}

func GetRoleItem(roleId int) (Role, bool) {
	var newRole Role
	result := Mysql.DB.Preload("Permission").First(&newRole, roleId)
	if result.Error != nil {
		return newRole, false
	}
	return newRole, true
}

func GetRoleList(name string, page int, pageSize int) map[string]interface{} {
	var roles []Role
	if name != "" {
		Mysql.DB.Where("name LIKE ?", name).Scopes(RationalData(), Mysql.Paginate(page, pageSize)).Preload("User").Find(&roles)
	} else {
		Mysql.DB.Scopes(RationalData(), Mysql.Paginate(page, pageSize)).Preload("User").Find(&roles)
	}

	var count int
	Mysql.DB.Table("role").Scopes(RationalData()).Count(&count)

	roleList := Paginator(page, pageSize, count)
	roleList["rows"] = roles

	return roleList
}

func DeleteRole(id string) (err error) {
	var role Role
	result := Mysql.DB.First(&role, id)
	if err = result.Error; err != nil {
		return errors.New("role not find")
	}
	//清除关联
	Mysql.DB.Model(&role).Association("Permission").Clear()

	//彻底删除
	Mysql.DB.Unscoped().Delete(&role)

	return nil
}

func BindPermissions(id int, perUids []string) error {
	var role Role
	var permissions []Permission

	Mysql.DB.Where("uid in (?)", perUids).Find(&permissions)

	Mysql.DB.First(&role, id)

	Mysql.DB.Model(&role).Association("Permission").Replace(permissions)

	return nil
}

func GetRolePermissions(id int) ([]string, error) {
	var role Role
	Mysql.DB.First(&role, id)

	var permissions []Permission
	Mysql.DB.Model(&role).Association("Permission").Find(&permissions)

	var uids []string
	for _, v := range permissions {
		uids = append(uids, v.Uid)
	}

	return uids, nil
}
