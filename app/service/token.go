package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId   int    `json:"userId"`
	PassWord string `json:"passWord"`
	jwt.StandardClaims
}

func CreateToken(userId int, password string) (string, error) {
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
