//go:generate protoc -I. -I"%GOPATH%/src" -I"%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2" --gofast_out=plugins=grpc:. "../../api/autoid.proto"

package server

import (
	"context"
	"errors"
	"golang.org/x/time/rate"
	"log"
	"orp/internal/app/myshop/api"
	"orp/internal/app/myshop/conf"
	"orp/internal/app/myshop/model"
	"orp/internal/app/myshop/service"
	"orp/pkg/consul"
	"time"
)

func New(svr *service.Service, c *conf.TomlConfig) {
	/*_, _, err := jaeger.NewTracer(serviceName, "192.168.99.101:6831")
	if err != nil {
		return
	}*/
	//defer i.Close()
	r, err := consul.NewRegister(c.App.CenterId, c.App.ServiceName, c.App.Port)
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
	api.RegisterShopServer(r.Server, &server{svr: svr, Limit: limit})
	r.Run()
}

type server struct {
	svr   *service.Service
	Limit *rate.Limiter
}

var _ api.ShopServer = &server{}

func (s *server) GetShopOrderLog(ctx context.Context, request *api.ShopRequest) (*api.ShopOrderResponse, error) {
	order := s.svr.GetShopOrderLog(request.Tid, request.RefundId)
	var orderLog  = api.ShopOrderResponse{Order: new(api.Order)}
	orderLog.Order.Id = order.Id
	orderLog.Order.Tid = order.Tid
	orderLog.Order.ShopId = order.ShopId
	orderLog.Order.RefundId = order.RefundId
	orderLog.Order.Type = int64(order.Type)
	orderLog.Order.CurrentVersion = int64(order.CurrentVersion)
	orderLog.Order.Content = order.Content
	return &orderLog, nil
}

func (s *server) AddShopOrderLog(ctx context.Context, orderLog *api.Order) (*api.ShopResponse, error) {
	var order model.TShopOrderLog
	order.ShopId = orderLog.ShopId
	order.RefundId = orderLog.RefundId
	order.Tid = orderLog.Tid
	order.Type = int8(orderLog.Type)
	order.CurrentVersion = uint8(orderLog.CurrentVersion)
	order.Content = orderLog.Content
	i, err := s.svr.AddShopOrderLog(order)
	response:=new(api.ShopResponse)
	if err != nil {
		response.Code = 1
		response.Msg = err.Error()
		return response, err
	}
	if i == 0 {
		response.Code = 1
		return response, errors.New("error")
	}
	return response, nil
}
