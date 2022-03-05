package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JWTsecret = []byte("somthing")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//GenerateToken 签发用户Token
func GenerateToken(id uint, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(7 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo-list",
		},
	}
	fmt.Println("claims", claims)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("tokenClaims", tokenClaims)
	token, err := tokenClaims.SignedString(JWTsecret)
	fmt.Println("token", token)
	fmt.Println("err", err)
	return token, err
}
func GenerateToken1(id uint, username string, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id:       id,
		UserName: username,
		// Authority: authority,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "to-do-list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

//ParseToken
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
