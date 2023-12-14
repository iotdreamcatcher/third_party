package sls

import (
	"fmt"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"time"
)

const (
	DebugLevel = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel
	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

type Level int8

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case DPanicLevel:
		return "dpanic"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

type Sls struct {
	Producer *producer.Producer
	Config   SlsConfig
}

type SlsConfig struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	ProjectName     string
	LogStoreName    string
}

type Template struct {
	AppName   string
	RequestId string
	Timestamp int64
	Message   string
	Method    string
	Path      string
	Data      string
}

func NewQuickSls(c SlsConfig, appName string) *QuickSls {
	Qs := new(QuickSls)
	s := new(Sls)
	s.Config = c
	p := InitProducer(c)
	s.Producer = p
	Qs.Sls = s
	Qs.PClient = p
	Qs.DefaultLevel = InfoLevel
	Qs.AppName = appName
	return Qs
}

func NewSls(c SlsConfig) *Sls {
	s := new(Sls)
	s.Config = c
	p := InitProducer(c)
	s.Producer = p
	return s
}

func InitProducer(config SlsConfig) *producer.Producer {
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = config.Endpoint
	producerConfig.AccessKeyID = config.AccessKeyID
	producerConfig.AccessKeySecret = config.AccessKeySecret
	producerInstance := producer.InitProducer(producerConfig)
	producerInstance.Start() // 启动producer实例
	return producerInstance
}

func (s Sls) Log(pclient *producer.Producer, ip string, level Level, template Template) error {
	sendData := map[string]string{
		"app_name":  template.AppName,
		"requestId": template.RequestId,
		"timestamp": time.Unix(template.Timestamp, 0).Format("2006-01-02 15:04:05"),
		"message":   template.Message,
		"method":    template.Method,
		"path":      template.Path,
		"data":      template.Data,
	}
	log := producer.GenerateLog(uint32(time.Now().Unix()), sendData)
	if err := pclient.SendLog(s.Config.ProjectName, s.Config.LogStoreName, level.String(), ip, log); err != nil {
		return err
	}
	return nil
}

type QuickSls struct {
	Sls          *Sls
	PClient      *producer.Producer
	AppName      string
	DefaultLevel Level
}

type QuickTemplate struct {
	RequestId string
	Message   string
	Method    string
	Path      string
	Data      string
}

func (qs QuickSls) Log(ip string, level Level, info QuickTemplate) error {
	sendData := map[string]string{
		"app_name":  qs.AppName,
		"requestId": info.RequestId,
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
		"message":   info.Message,
		"method":    info.Method,
		"path":      info.Path,
		"data":      info.Data,
	}
	log := producer.GenerateLog(uint32(time.Now().Unix()), sendData)
	if err := qs.PClient.SendLog(qs.Sls.Config.ProjectName, qs.Sls.Config.LogStoreName, level.String(), ip, log); err != nil {
		return err
	}
	return nil
}
