package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"golang.org/x/time/rate"
	"log"
	"orp/internal/app/myapp/api"
	service2 "orp/internal/app/myapp/service"
	"orp/pkg/consul"
	"orp/pkg/jaeger"
	"time"
)

const serviceName = "HelloService"

func main() {
	_, i, err := jaeger.NewTracer(serviceName, "192.168.99.101:6831")
	if err != nil {
		return
	}
	defer i.Close()
	r, err := consul.NewRegister("192.168.99.101:8500", serviceName, 8080)
	if err != nil {
		log.Fatal(err)
	}
	//创建限流器 初始容量为10，每秒产生一个令牌
	limit := rate.NewLimiter(rate.Every(time.Second), 15)

	//创建熔断器
	hystrix.ConfigureCommand("HelloService", hystrix.CommandConfig{
		Timeout:                2000, //超时时间设置  单位毫秒
		MaxConcurrentRequests:  8,    //最大请求数
		SleepWindow:            1,    //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  30,   //错误率
		RequestVolumeThreshold: 5,    //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	})
	api.RegisterOrderServiceServer(r.Server, &service2.Order{Limit: limit})
	r.Run()
}
