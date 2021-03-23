package Models

import (
	Mysql "SimpleApi/databses"
	Middlewares "SimpleApi/middleware"
	jwt "SimpleApi/pkg/utils"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(userData UserForm) (uint, error) {

	var UserList []User
	//查重
	err := Mysql.DB.Table("user").Where("username = ?", userData.Username).Find(&UserList).Error
	if err != nil {
		return 0, errors.New("user search fail")
	} else {
		if len(UserList) > 0 {
			return 0, errors.New("user exit")
		}
	}

	originPwd := []byte(userData.User.Password)
	//此方法生成hash值
	hashPwd, _ := bcrypt.GenerateFromPassword(originPwd, bcrypt.DefaultCost) //password为string类型
	userData.User.Password = string(hashPwd)
	result := Mysql.DB.Create(&userData.User)
	id := userData.User.ID

	var roles []Role
	Mysql.DB.Find(&roles, userData.Roles)

	//更新关系
	Mysql.DB.Model(&userData.User).Association("Role").Replace(roles)

	if result.Error != nil {
		return 0, errors.New("user create fail")
	}
	return id, nil
}

//过滤root
func NoRoot() func(db *gorm.DB) *gorm.DB {
	return func(DB *gorm.DB) *gorm.DB {
		return DB.Where("type != ?", "00")
	}
}

func GetUserList(username string, page int, pageSize int) map[string]interface{} {
	var users []User
	if username != "" {
		Mysql.DB.Where("username = ?", username).Scopes(Mysql.Paginate(page, pageSize), NoRoot()).Preload("Role").Order("created_at desc").Find(&users)
	} else {
		Mysql.DB.Scopes(Mysql.Paginate(page, pageSize), NoRoot()).Preload("Role").Order("created_at desc").Find(&users)
	}

	var count int
	Mysql.DB.Table("user").Count(&count)

	userList := Paginator(page, pageSize, count)
	userList["rows"] = users

	return userList
}

//返回用户及权限
type UserAndPermission struct {
	User        User     `json:"user"`
	Permissions []string `json:"permissions"`
}

func GetUserItem(userId int) (user UserForm, err error) {
	var userInfo UserForm
	result := Mysql.DB.Preload("Role").First(&userInfo.User, userId)

	var roleIds []int
	for i := 0; i < len(userInfo.Role); i++ {
		roleIds = append(roleIds, int(userInfo.Role[i].ID))
	}

	userInfo.Roles = roleIds

	if result.Error != nil {
		return userInfo, errors.New("user not find")
	}
	return userInfo, nil
}

func GetLoginUser() (UserAndPermission, bool) {

	j := jwt.NewJWT()
	jwtClaims, _ := j.ParseToken(Middlewares.AuthToken)

	var resultData UserAndPermission

	result := Mysql.DB.Preload("Role").First(&resultData.User, jwtClaims.Id)
	if result.Error != nil {
		return resultData, false
	}

	if resultData.User.Type == "00" {
		var permissions []Permission
		Mysql.DB.Find(&permissions)
		for _, value := range permissions {
			resultData.Permissions = append(resultData.Permissions, value.Uid) // 追加1个元素
		}
	} else {
		var roles []Role
		Mysql.DB.Model(&resultData.User).Association("Role").Find(&roles)

		for _, item := range roles { //item是值拷贝
			var permissions []Permission
			Mysql.DB.Model(&item).Association("Permission").Find(&permissions)
			for _, value := range permissions {
				resultData.Permissions = append(resultData.Permissions, value.Uid) // 追加1个元素
			}
		}
		resultData.Permissions = removeDuplicateElement(resultData.Permissions)
	}
	return resultData, true
}

//数组去重
func removeDuplicateElement(data []string) []string {
	result := make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func BindRoles(id int, roles []Role) (err error) {
	var user User

	result := Mysql.DB.First(&user, id)
	if err = result.Error; err != nil {
		return errors.New("user not find")
	}

	Mysql.DB.Model(&user).Association("Role").Replace(roles)

	return nil
}

func EditUser(userId int, formData UserForm) bool {

	var user User
	var roles []Role
	Mysql.DB.First(&user, userId)

	//更新数据
	Mysql.DB.Model(&user).Update(formData.User)

	Mysql.DB.Find(&roles, formData.Roles)

	//更新关系
	Mysql.DB.Model(&user).Association("Role").Replace(roles)

	return true
}

func DeleteUser(id int) (err error) {
	var user User
	result := Mysql.DB.First(&user, id)
	if err = result.Error; err != nil {
		return errors.New("user not find")
	}

	//清除关联
	Mysql.DB.Model(&user).Association("Role").Clear()

	//删除
	Mysql.DB.Unscoped().Delete(&user)

	return nil
}
