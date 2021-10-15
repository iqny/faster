package logger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

func NewLfsHook() logrus.Hook {
	debugWriter, err := setRotateLogs("debug")
	infoWriter, err := setRotateLogs("info")
	errorWriter, err := setRotateLogs("error")
	warnWriter, err := setRotateLogs("warn")
	traceWriter, err := setRotateLogs("trace")
	fatalWriter, err := setRotateLogs("fatal")
	panicWriter, err := setRotateLogs("panic")

	if err != nil {
		logrus.Errorf("config local file system for logger error: %v", err)
	}
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: debugWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: fatalWriter,
		logrus.PanicLevel: panicWriter,
		logrus.TraceLevel: traceWriter,
	}, &LogFormatter{})

	return lfsHook
}
func setRotateLogs(level string) (writer *rotatelogs.RotateLogs, err error) {
	writer, err = rotatelogs.New(
		strings.Join([]string{c.FileName,"/",c.AppName,"/", level, ".%Y%m%d", ".log"}, ""),
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		//rotatelogs.WithLinkName(fileName),

		// WithRotationTime设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(12),
	)
	return
}
