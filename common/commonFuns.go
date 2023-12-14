package common

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/peer"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author joker
 * @Description //TODO 常用函数库
 * @Date 2020-7-12 17:08:54
 **/

// 注意client 本身是连接池，不要每次请求时创建client
var (
	HttpClient = &http.Client{
		Timeout: 30 * time.Second,
	}
)

func BindAndCheck(ctx *gin.Context, data interface{}) error {
	if err := ctx.ShouldBindJSON(data); err != nil {
		return errors.New(fmt.Sprintf("bindjson err%s", err))
	}
	// 校验数据
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return errors.New(fmt.Sprintf("validator err%s", err))
	}
	return nil
}

func RandInt64(min, max int64) int {
	if min >= max || min == 0 || max == 0 {
		return int(max)
	}
	return int(rand.Int63n(max-min) + min)
}

func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return int(max)
	}
	return rand.Intn(max-min) + min
}

func DelFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return nil //f为空 错误不为空 错误是文件不存在 可以忽略
		}
		if f.IsDir() {
			fmt.Printf("文件夹 继续递归 %s  \n", path)
			DelFilelist(path)
		} else {
			err := os.RemoveAll(path)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("删除文件 %s  \n", path)
			return nil
		}
		//	println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("walk 错误 err: %v\n", err)
	}
}

// RandStringRunes 返回随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func UploadFile(url string, params map[string]string, nameField, fileName string, file io.Reader) ([]byte, error) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func DistributeFile(url string, params map[string]string, nameField, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)

	formFile, err := writer.CreateFormFile(nameField, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(formFile, file)
	if err != nil {
		return nil, err
	}

	for key, val := range params {
		writer.WriteField(key, val)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	//req.Header.Set("Content-Type","multipart/form-data")
	req.Header.Add("Content-Type", writer.FormDataContentType())

	resp, err := HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return content, nil
}

// Replace 根据替换表执行批量替换
func Replace(table map[string]string, s string) string {
	for key, value := range table {
		s = strings.Replace(s, key, value, -1)
	}
	return s
}

// Int2Str int类型转string类型
func Int2Str(inter int) string {
	string := strconv.Itoa(inter)
	return string
}

// Int2Str int64类型转string类型 精确到后2位 精确到后4位
func Int642Str(inter int64) string {

	string := strconv.FormatInt(inter, 10)
	return string
}

func Int64Str(inter int64) string {
	string := strconv.FormatInt(inter, 10)
	return string
}

func Str2Float64(in string) float64 {
	//num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	float, _ := strconv.ParseFloat(in, 64)
	return float
}

// Str2Int string类型转Int类型
func Str2Int(inter string) int {
	int, err := strconv.Atoi(inter)
	if err != nil {
		log.Println("err", err)
	}
	return int
}

// Str2Int string类型转Int类型
func Str2Uint(inter string) uint {
	uint64, _ := strconv.ParseUint(inter, 10, 64)
	return uint(uint64)
}

// Str2Int64 string类型转Int64类型
func Str2Int64(inter string) int64 {
	int64, _ := strconv.ParseInt(inter, 10, 64)
	return int64
}

func Arr2Str(strings []string) string {
	b, _ := json.Marshal(strings)
	return fmt.Sprintf("%s", b)
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func GetClientIP(ctx context.Context) (string, error) {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("[getClinetIP] invoke FromContext() failed")
	}
	if pr.Addr == net.Addr(nil) {
		return "", fmt.Errorf("[getClientIP] peer.Addr is nil")
	}
	addSlice := strings.Split(pr.Addr.String(), ":")
	return addSlice[0], nil
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// TODO: 获取当月的最后第一天或者最后一天
func ReturnSpecifyMonth(year, month int) (time.Time, time.Time) {
	//currentYear, currentMonth, _ := now.Date()
	now := time.Now()
	currentLocation := now.Location()

	firstOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, 0)
	return firstOfMonth, lastOfMonth
}

// TODO: 获取当年的最后第一天或者最后一天
func ReturnSpecifyYear(year int) (time.Time, time.Time) {
	//currentYear, currentMonth, _ := now.Date()
	now := time.Now()
	currentLocation := now.Location()

	firstOfMonth := time.Date(year, 1, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 12, 0)
	return firstOfMonth, lastOfMonth
}
