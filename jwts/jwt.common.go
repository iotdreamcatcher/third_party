package jwts

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtCommonClaims struct {
	UserId    uint
	LoginTime time.Time
	jwt.StandardClaims
}

func JwtCommonCreateToken(exp int, UserId uint, key string, loginTime time.Time) (string, int64, error) {
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
func JwtCommonParseToken(tokenString, key string) (*JwtWithoutDomainClaims, error) {
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
