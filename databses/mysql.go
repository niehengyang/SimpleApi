package Mysql

import (
	"SimpleApi/pkg/setting"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	connArgs := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DB_USER, setting.DB_PWD, setting.DB_HOST, setting.DB_PORT, setting.DB_NAME)

	DB, err = gorm.Open(setting.DB_TYPE, connArgs)
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
		return nil
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
		return nil
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	fmt.Println("数据库连接成功！")
	return DB
}

//分页封装
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(DB *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return DB.Offset(offset).Limit(pageSize)
	}
}
