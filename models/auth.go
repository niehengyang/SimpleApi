package Models

import (
	Mysql "SimpleApi/databses"
	jwt "SimpleApi/pkg/utils"
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginInput struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type JwtToken struct {
	User   interface{} `json:"user"`
	Token  string      `json:"token"`
	Expire int64       `json:"expire"`
}

func Login(username string, password string) (token JwtToken, err error) {
	var user User
	var nullData JwtToken

	obj := Mysql.DB.Where("username = ?", username).First(&user)
	if err = obj.Error; err != nil {
		return
	}

	//验证密码
	checkResult := ComparePasswords(user.Password, []byte(password))
	if !checkResult {
		return nullData, errors.New("invalid password")
	}

	generateToken := GenerateToken(user)

	user.Token = generateToken.Token
	user.Expire = generateToken.Expire

	Mysql.DB.Save(&user)

	return generateToken, nil
}

// 验证密码
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}

// 生成令牌  创建jwt风格的token
func GenerateToken(user User) JwtToken {

	j := jwt.NewJWT()

	claims := jwt.CustomClaims{
		user.ID,
		user.Username,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 7200), // 过期时间 两小时
			Issuer:    "nhy",                           //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return JwtToken{
			User:   user,
			Token:  token,
			Expire: int64(time.Now().Unix() + 3600),
		}
	}

	data := JwtToken{
		User:   user,
		Token:  token,
		Expire: int64(time.Now().Unix() + 3600),
	}
	return data
}
