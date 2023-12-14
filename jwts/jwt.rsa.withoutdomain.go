package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtWithoutDomainClaims struct {
	UserId    uint
	LoginTime time.Time
	jwt.StandardClaims
}

func JwtWithoutDomainCreateToken(exp int, UserId uint, privateKey string, loginTime time.Time) (string, int64, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", -1, err
	}
	expiresAt := time.Now().Add(time.Duration(exp) * time.Second).Unix()
	customClaims := &JwtWithoutDomainClaims{
		UserId:    UserId,
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
func JwtWithoutDomainParseToken(tokenString, pubKey string) (*JwtWithoutDomainClaims, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &JwtWithoutDomainClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if claims, ok := token.Claims.(*JwtWithoutDomainClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
