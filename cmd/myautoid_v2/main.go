package main

import (
	"flag"
	"orp/internal/app/myautoid_v2/conf"
	grpc "orp/internal/app/myautoid_v2/server/grpc"
	"orp/internal/app/myautoid_v2/service"
)

var s *service.Service

func main() {
	flag.Parse()
	c:=conf.Config()
	s = service.New(c)
	//log.Init(conf.Cfg.Log)
	/*loger.Init(conf.Cfg.Loger, func() logrus.Hook {
		return loger.NewLfsHook()
	})*/
	defer s.Close()
	grpc.New(s,c)
}
