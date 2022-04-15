package utils

import (
	"JeffMusic/setting"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("i.am.a.sheep")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

func CreateToken(id int) (string, error) {
	var err error
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(), //颁发时间
			Issuer:    setting.Conf.Host, // 签名颁发者
			Subject:   "user token",      //签名主题
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

// ValidateToken 验证token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, _, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return nil, err
	}
	return token, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}
