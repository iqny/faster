package main

import (
	"flag"
	"log"
	"orp/internal/app/myapp/conf"
	"orp/internal/app/myapp/http"
	grpc "orp/internal/app/myapp/server/grpc"
	"orp/internal/app/myapp/service"
	"time"
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
	grpc.New(s)
	r := http.Init("c", s)
	time.Local, _ = time.LoadLocation(c.App.Timezone)
	/*server := endless.NewServer(conf.Cfg.App.Host, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()*/
	err:=r.Run(c.App.Host)
	if err != nil {
		log.Fatalf("Server err: %v", err)
	}
}
