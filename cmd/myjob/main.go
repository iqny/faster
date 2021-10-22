//go:generate goversioninfo -icon=../../resource/icon.ico -manifest=../../resource/goversioninfo.exe.manifest
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"orp/internal/app/myjob/conf"
	"orp/internal/app/myjob/job"
	"orp/pkg/rabbitmq"
	"orp/pkg/rabbitmq/logger"
	"os"
	"os/signal"
	"syscall"
)

var q *rabbitmq.Queue

func main() {
	flag.Parse()
	ctx, cancelFunc := context.WithCancel(context.Background())
	conf.Config()
	logger.Init(conf.Cfg.Logger, func() logrus.Hook {
		return logger.NewLfsHook()
	})
	q = rabbitmq.New(ctx, conf.Cfg.Queue)

	j:=job.Job{}
	q.Register("testJob",j.TestJob)
	q.Register("orderJob",j.OrderJob)
	q.Register("orderTransformJob",j.OrderTransformJob)
	q.Run()
	signalHandler(cancelFunc)
}

func signalHandler(cancelFunc context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case sig := <-ch:
			fmt.Printf("get a signal %s, stop the apm-queue process\n", sig.String())
			switch sig {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				cancelFunc()
				q.Wait()
				fmt.Println("all obj work quit.")
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}
}
