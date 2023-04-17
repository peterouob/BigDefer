package utils

import (
	"bigdefer/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 加密token
func SignToken(token *Token) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, token).
		SignedString([]byte(config.Config.GetString("jwt.key")))
}

// 解密token
func ParseToken(sign string) (*Token, error) {
	tokenClaims, err := jwt.ParseWithClaims(sign, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Config.GetString("jwt.key")), nil
	})
	if err != nil {
		return nil, err
	} else {
		if tokenClaims == nil {
			return nil, errors.New("token claim error")
		} else {
			if claims, ok := tokenClaims.Claims.(*Token); ok {
				return claims, nil
			} else {
				return nil, errors.New("token claim error")
			}
		}
	}
}
