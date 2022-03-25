package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/sirupsen/logrus"
	"orp/pkg/rabbitmq/logger"
	"testing"
	"time"
)

func TestRun(t *testing.T) {

	tm,_:=time.ParseDuration("4h")
	fmt.Println(time.Duration(tm))
	return
	c:=&logger.Config{
		FileName:         "D:/juest/src/orp/logs",
		FileNameFormat:   ".%Y%m%d",
		WithRotationTime: 1,
		LogFileExpired:   1,
		ReportCaller:     true,
		ShowConsole:      false,
		Level:            "debug",
		AppName:          "queue",
	}
	logger.Init(c, func() logrus.Hook {
		return logger.NewLfsHook()
	})
	//j:=job.Job{}
	for i := 0; i < 10; i++ {
		var scb string
		b,_:=byteEncoder("abc")
		bytes.NewReader(b)
		decoder := gob.NewDecoder(bytes.NewReader(b))
		decoder.Decode(&scb)
	}
}
func byteEncoder(s interface{}) ([]byte,error) {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return result.Bytes(),nil
}
