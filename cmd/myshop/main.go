package main

import (
	"flag"
	"orp/internal/app/myshop/conf"
	grpc "orp/internal/app/myshop/server/grpc"
	"orp/internal/app/myshop/service"
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
