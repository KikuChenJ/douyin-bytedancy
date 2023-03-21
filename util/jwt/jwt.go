package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/hjk-cloud/douyin/models"
	"time"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

//定义Secret
var mySecret = []byte("douyin")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

//定义JWT的过期时间
const TokenExpireDuration = time.Hour * 6

//GenToken 生成 Token
func GenToken(userId int) (Token string, err error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		userId, // 自定义字段
		jwt.StandardClaims{ // JWT规定的7个官方字段
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "badman",                                   // 签发人
		},
	}
	// 加密并获得完整的编码后的字符串token
	Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	return Token, err
}

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	// 解析token
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid { // 校验token
		err = errors.New("invalid token")
	}
	return
}

// JWTAuth 用于验证token，并返回token对应的username
func JWTAuth(token string) (int, error) {
	if token == "" {
		return 0, errors.New("token为空")
	}
	claim, err := ParseToken(token)
	if err != nil {
		return 0, errors.New("token过期")
	}
	//最后验证这个user是否真的存在
	if _, err := models.NewUserDaoInstance().QueryUserById(claim.UserId); err != nil {
		return 0, errors.New("user不存在")
	}

	return claim.UserId, nil
}
