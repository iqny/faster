package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"testing"
)

func TestNewLogger(t*testing.T){
	c:=&Config{
		FileName:         "D:/juest/src/orp/logs",
		FileNameFormat:   ".%Y%m%d",
		WithRotationTime: 1,
		LogFileExpired:   1,
		ReportCaller:     true,
		ShowConsole:      false,
		Level:            "debug",
		AppName:          "queue",
	}
	Init(c,func() logrus.Hook {
		return NewLfsHook()
	})
	for i:=1;i<1000000;i++ {
		fields := make(map[string]interface{})
		fields["uuid"] = "abrr"
		fields["appName"] = "que"
		fields["file"] = fileInfo(2)
		Debug(fields, "run-----")
	}
}
func fileInfo(skip int) string {
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
