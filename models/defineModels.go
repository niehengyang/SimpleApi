package Models

import (
	Middlewares "SimpleApi/middleware"
	jwt "SimpleApi/pkg/utils"
	"database/sql/driver"
	"fmt"
	"github.com/jinzhu/gorm"
	"math"
	"time"
)

// 1. 创建 time.Time 类型的副本 XTime；
type XTime struct {
	time.Time
}

const TimeFormat = "2006-01-02 15:04:05"

//MyTime 自定义时间
type MyTime time.Time

func (t *XTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = XTime{now}
	return err
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 5. 自定义 BaseModel，结构和 gorm.Model 一致，将 time.Time 替换为 LocalTime；
type BaseModel struct {
	ID        uint `gorm:"primary_key,AUTO_INCREMENT"`
	CreatedAt XTime
	UpdatedAt XTime
	DeletedAt *XTime `sql:"index"`
}

// 用户
type User struct {
	BaseModel
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Username string `gorm:"size:100;column(username)" json:"username" form:"username"`
	Password string `gorm:"size:255;column(password)" json:"password" form:"password"`
	Email    string `gorm:"size:200;column(email)" json:"email" form:"email"`
	Status   string `gorm:"size:20;column(status)" json:"status" form:"status"`
	Avatar   string `gorm:"size:200;column(avatar)" json:"avatar" form:"avatar"`
	Type     string `gorm:"size:20;column(type)" json:"type" form:"type"`
	Phone    string `gorm:"size:20;column(phone)" json:"phone" form:"phone"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
	Token    string `gorm:"size:500;column:token" json:"-"`
	Expire   int64  `gorm:"column:expire" json:"-"`
	Role     []Role `gorm:"many2many:user_roles"`
}

//权限
type Permission struct {
	BaseModel
	Name     string `gorm:"size:100;column(name)" json:"name" form:"name"`
	Parent   string `gorm:"size:10;column(parent)" json:"parent" form:"parent"`
	Status   string `gorm:"size:10;column(status)" json:"status" form:"status"`
	Uid      string `gorm:"size:50;column(uid)" json:"uid" form:"uid"`
	Type     string `gorm:"size:10;column(type)" json:"type" form:"type"`
	Url      string `gorm:"size:500;column(url)" json:"url" form:"url"`
	Icon     string `gorm:"size:500;column(icon)" json:"icon" form:"icon"`
	Describe string `gorm:"size:200;column(describe)" json:"describe" form:"describe"`
}

//角色
type Role struct {
	BaseModel
	Name       string       `gorm:"size:100;column(name)" json:"name" form:"name"`
	UserId     int          `gorm:"" json:"user_id"`
	Desc       string       `gorm:"size:500;column(desc)" json:"desc" form:"desc"`
	User       User         `gorm:"foreignKey:UserId"`
	Permission []Permission `gorm:"many2many:role_permissions"`
}

func StartMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Permission{}, &Role{})
	return
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page int, prepage int, nums int) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["total"] = nums
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}

//合法数据(当前登录用户)
func RationalData() func(db *gorm.DB) *gorm.DB {

	j := jwt.NewJWT()
	JwtCustomClaims, _ := j.ParseToken(Middlewares.AuthToken)

	if JwtCustomClaims.Id == 1 {
		return func(DB *gorm.DB) *gorm.DB {
			return DB
		}
	} else {
		constId := jwt.IntToString(int(JwtCustomClaims.Id))
		return func(DB *gorm.DB) *gorm.DB {
			return DB.Where("user_id = ?", constId)
		}
	}
}
