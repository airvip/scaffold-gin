package util

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)


var jwtSecret = []byte("hello")

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

type Claims struct {
	Id string
	jwt.StandardClaims
}

func ReleaseToken(userId, nickName string) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			Audience:  nickName,              // 受众
			ExpiresAt: expirationTime.Unix(), // 失效时间
			Id:        userId,                //编号
			IssuedAt:  time.Now().Unix(),     // 签发时间
			Issuer:    "airvip",              //签发人
			NotBefore: time.Now().Unix(),     //生效时间
			Subject:   "user token",          //主题
		},
	}
	/* claims := jwt.StandardClaims{
		Audience:  user.Nickname,         // 受众
		ExpiresAt: expirationTime.Unix(), // 失效时间
		Id:        user.Identity,         //编号
		IssuedAt:  time.Now().Unix(),     // 签发时间
		Issuer:    "airvip",              //签发人
		NotBefore: time.Now().Unix(),     //生效时间
		Subject:   "user token",          //主题
	} */

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenString, nil

}

// 解析 token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

