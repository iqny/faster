package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func newLog(filePath string) *logrus.Logger {
	logger := logrus.New()
	if filePath != "" {
		writer, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.Fatalf("create file log.txt failed: %v", err)
		}

		//logger.SetOutput(io.MultiWriter(writer))
		logger.SetOutput(writer)
	}
	return logger
}
