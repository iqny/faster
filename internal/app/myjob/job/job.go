package job

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
)

type Job struct {
	uuid    string
	appName string
}

func (j *Job) setInfo(taskId string, name string) {
	j.uuid = taskId
	j.appName = name
}
func (j *Job) debug(args ...interface{}) {
	fields := make(map[string]interface{})
	fields["uuid"] = j.uuid
	fields["appName"] = j.appName
	fields["file"] = fileInfo(2)
	logrus.WithFields(fields).Debug(args...)
}
func (j *Job) info(args ...interface{}) {
	fields := make(map[string]interface{})
	fields["uuid"] = j.uuid
	fields["appName"] = j.appName
	fields["file"] = fileInfo(2)
	logrus.WithFields(fields).Info(args...)
}
func (j *Job) warn(args ...interface{}) {
	fields := make(map[string]interface{})
	fields["uuid"] = j.uuid
	fields["appName"] = j.appName
	fields["file"] = fileInfo(2)
	logrus.WithFields(fields).Warn(args...)
}
func (j *Job) error(args ...interface{}) {
	fields := make(map[string]interface{})
	fields["uuid"] = j.uuid
	fields["appName"] = j.appName
	fields["file"] = fileInfo(2)
	logrus.WithFields(fields).Error(args...)
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
