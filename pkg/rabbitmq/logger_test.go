package rabbitmq

import "testing"

func TestNewLogger(t *testing.T){
	s:="D:/juest/src/orp/logs/queue/monitor1.log"
	log:=newLog(s)
	for i:=0;i<100000;i++ {
		log.Infoln("abc1")
	}
}
