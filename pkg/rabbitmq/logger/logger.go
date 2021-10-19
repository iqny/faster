package logger

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	FileName         string //文件名
	FileNameFormat   string //文件时间后缀
	WithRotationTime int    //按每天分割文件
	LogFileExpired   int    //保存多久文件，1天单位
	ReportCaller     bool   //是否显示文件名及行号
	ShowConsole      bool   //控制台是否显示
	Level            string //级别
	AppName          string
}

var c *Config

type FunHook func() logrus.Hook

func Init(conf *Config, hook FunHook) {
	setNull()
	c = conf
	if false == conf.ShowConsole {
		setNull()
	}
	logrus.SetReportCaller(c.ReportCaller)
	setLevel(c.Level)
	if hook != nil {
		logrus.AddHook(hook())
	}
}

/*var appName string

func SetAppName(name string) {
	appName = name
}*/
func setLevel(level string) {
	switch level {
	//如果日志级别不是debug就不要打印日志到控制台了
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
		//logrus.SetOutput(os.Stderr)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetReportCaller(true)
	case "error":
		logrus.SetReportCaller(true)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var file string
	//var line int
	/*if entry.Caller != nil {
		file = entry.Caller.File
		line = entry.Caller.Line
	}*/
	//level := strings.ToUpper(entry.Level.String())
	var uuid string
	if id, ok := entry.Data["uuid"]; ok {
		uuid = id.(string)
		delete(entry.Data, "uuid")
	}
	if f, ok := entry.Data["file"]; ok {
		file = f.(string)
		delete(entry.Data, "file")
	}
	/*if name, ok := entry.Data["appName"]; ok {
		appName = name.(string)
		delete(entry.Data, "appName")
	}*/
	content, _ := json.Marshal(entry.Data)
	msg := fmt.Sprintf("%s [%s] [GOID:%d] [%s]  #msg:%s #content:%s \n", timestamp, uuid, getGID(), file, entry.Message, string(content))
	//msg := fmt.Sprintf("%s [%s] [%s:%d] %s #msg:%s #content:%s \n", timestamp, uuid, file, line, appName, entry.Message, string(content))

	return []byte(msg), nil
}

// 获取当前协程id
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
func setNull() {
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	writer := bufio.NewWriter(src)
	logrus.SetOutput(writer)
}
func CreateUUid() string {
	s := strings.Split(uuid.NewV4().String(), "-")
	return s[0]
}
func FileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
