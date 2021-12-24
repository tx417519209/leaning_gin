package util

import (
	"time"

	"github.com/golang-jwt/jwt"
	"xing.learning.gin/src/gin-blog/pkg/setting"
)

var jwtSecret = setting.JwtSecret
var issuer = setting.JwtIssuer

type CustomClaims struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := CustomClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtSecret))
	return token, err
}

func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{},
		func(*jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
