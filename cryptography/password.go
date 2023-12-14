package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"regexp"
)

func Password(password string, salt string) string {
	sha := sha256.New()
	sha.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(sha.Sum([]byte(salt)))
}

func CheckPassword(password string) bool {
	if ok, _ := regexp.MatchString("^[a-z0-9A-Z]{8,20}$", password); !ok {
		return false
	}
	if ok, _ := regexp.MatchString("[A-Z]{1,20}", password); !ok {
		return false
	}
	return true
}

func CheckPasswordAes(password, salt, key, baseCode string) bool {
	//解析
	original := AesDecryptByCTR(baseCode, key)
	if (password + salt) == original {
		return true
	}
	return false
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandSalt(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
