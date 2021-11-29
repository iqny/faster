//go:generate protoc -I. -I"%GOPATH%/src" -I"%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2" --gofast_out=plugins=grpc:. "../../api/autoid.proto"

package server

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"orp/internal/app/myautoid/api"
	"orp/internal/app/myautoid/conf"
	"orp/internal/app/myautoid/service"
	"orp/pkg/consul"
	"time"
)

func New(svr *service.Service,c *conf.TomlConfig) {
	/*_, _, err := jaeger.NewTracer(serviceName, "192.168.99.101:6831")
	if err != nil {
		return
	}*/
	//defer i.Close()
	r, err := consul.NewRegister(c.App.CenterId, c.App.ServiceName,c.App.Port )
	if err != nil {
		log.Fatal(err)
	}
	//创建限流器 初始容量为10，每秒产生一个令牌
	limit := rate.NewLimiter(rate.Every(time.Second), 15)

	/*//创建熔断器
	hystrix.ConfigureCommand(serviceName, hystrix.CommandConfig{
		Timeout:                2000, //超时时间设置  单位毫秒
		MaxConcurrentRequests:  8,    //最大请求数
		SleepWindow:            1,    //过多长时间，熔断器再次检测是否开启。单位毫秒
		ErrorPercentThreshold:  30,   //错误率
		RequestVolumeThreshold: 5,    //请求阈值  熔断器是否打开首先要满足这个条件；这里的设置表示至少有5个请求才进行ErrorPercentThreshold错误百分比计算
	})*/
	api.RegisterAutoIdServer(r.Server, &server{svr: svr, Limit: limit})
	r.Run()
}

type server struct {
	svr   *service.Service
	Limit *rate.Limiter
}

func (s *server) Get(ctx context.Context, order *api.Request) (*api.Response, error) {
	id, err := s.svr.GetID(int16(order.Code))
	if err != nil {
		return &api.Response{
			Id: 0,
			Msg:  err.Error(),
		}, nil
	}
	return &api.Response{
		Id: id,
		Msg:  "get id success",
	}, nil
}

var _ api.AutoIdServer = &server{}

