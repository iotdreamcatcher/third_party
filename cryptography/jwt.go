/*
*

	@author: taco
	@Date: 2023/7/26
	@Time: 17:47

*
*/
package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	UserId    uint
	TenantID  uint
	LoginTime time.Time
	jwt.StandardClaims
}

func CreateToken(exp int, UserId, TenantID uint, privateKey string, loginTime time.Time) (string, int64, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", -1, err
	}
	expiresAt := time.Now().Add(time.Duration(exp) * time.Second).Unix()
	customClaims := &CustomClaims{
		UserId:    UserId,
		TenantID:  TenantID,
		LoginTime: loginTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt, // 过期时间
		},
	}
	//采用 RS256 加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, customClaims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", -1, err
	}
	return tokenString, expiresAt, nil
}

// 解析 token
func ParseToken(tokenString, pubKey string) (*CustomClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
