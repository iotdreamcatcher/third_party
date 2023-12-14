/*
*

	@author: taco
	@Date: 2023/7/21
	@Time: 13:44

*
*/
package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func ToUp(sign string) string {
	arr := strings.Split(sign, "")
	tempS := ""
	for _, i2 := range arr {
		if IsNumber(i2) {
			tempS += " " + i2 + " "
		} else {
			tempS += i2
		}
	}
	arr2 := strings.Split(tempS, " ")
	tempB := ""
	for _, k := range arr2 {
		if k == " " {
			continue
		} else if IsNumber(k) {
			tempB += k
		} else {
			tempB += strings.Title(k)
		}
	}
	return tempB
}

func IsNumber(str string) bool {
	pattern := "^[0-9]+$"
	match, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	return match
}
func Random() (string, error) {
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	uuidStr := uuid.NewString()
	arr := strings.Split(uuidStr, "-")
	var temp string
	for _, s := range arr {
		temp += s
	}
	arr2 := strings.Split(temp, "")
	var tempS string
	var tempR string
	for i, s := range arr2 {
		if (i+1)%2 == 0 {
			tempS += s
			s1, err := strconv.ParseInt(tempS, 16, 0)
			if err != nil {
				return tempR, err
			}
			tempR += chars[s1%0x24]
			tempS = ""
		} else {
			tempS += s
		}
	}
	return tempR, nil
}

func Md5SumWithString(hash string, salt string) string {
	h := md5.New()
	h.Write([]byte(hash + salt))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateTicket(pfx string, l int) string {
	rand.Seed(time.Now().UnixNano())
	var TicketRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	t := make([]rune, l)
	for i := range t {
		t[i] = TicketRunes[rand.Intn(len(TicketRunes))]
	}
	if pfx == "" {
		return fmt.Sprintf("%s", string(t))
	}
	return fmt.Sprintf("%s-%s", pfx, string(t))
}
