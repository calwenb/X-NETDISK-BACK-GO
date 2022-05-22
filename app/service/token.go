package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"netdisk-back/app/util"
	"time"
)

type Claims struct {
	UserId   string `json:"userId"`
	PassWord string `json:"passWord"`
	jwt.StandardClaims
}

var jwtSecret = []byte("wen")

const TOKEN_PREFIX = "token:"

func CreateToken(userId string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * 30 * time.Hour)
	issuer := "Mr.文"
	claims := Claims{
		UserId:   userId,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

/*
1.redis上是否有令牌
2.令牌的权限是否为管理员
*/
func VerifyToken(token string) bool {
	rdb := util.GetRDBConn()
	value, err := rdb.Get(TOKEN_PREFIX + token).Result()
	if err != nil {
		err.Error()
		return false
	}
	if value == "1" {
		return true
	}
	return false
}
func ParseToken(token string) (Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return Claims{}, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return Claims{}, nil
}

func ParseToken2(tokens string) (*Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokens, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("golang"), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	return Claims, nil
}

func ParseToken3(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Printf("", claims.Audience)
			return claims, nil
		}
	}

	return nil, err

}
