//go:generate protoc -I. -I"%GOPATH%/src" -I"%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2" --gofast_out=. "../../api/order.proto"

package server

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"orp/internal/app/myapp/api"
	"orp/internal/app/myapp/service"
	"orp/pkg/consul"
	"time"
)

const serviceName = "HelloService"

func New(svr *service.Service) {
	/*_, _, err := jaeger.NewTracer(serviceName, "192.168.99.101:6831")
	if err != nil {
		return
	}*/
	//defer i.Close()
	r, err := consul.NewRegister("127.0.0.1:8500", serviceName, 8080)
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
	api.RegisterOrderServiceServer(r.Server, &server{svr: svr, Limit: limit})
	go r.Run()
}

type server struct {
	svr   *service.Service
	Limit *rate.Limiter
}

func (s *server) Add(ctx context.Context, order *api.Order) (*api.Response, error) {
	//addOrderSpan,_ := opentracing.StartSpanFromContext(ctx, "AddOrder")
	//defer addOrderSpan.Finish()
	/*if allow := s.Limit.Allow(); !allow {
		addOrderSpan.SetTag("order.Limit.Allow()", false)
		//log.Println("limit cancel")
		return nil, errors.New("limit cancel")
	}*/
	//addOrderSpan.SetTag("add.order start", true)
	_, err := s.svr.Add(order.Name)
	if err != nil {
		return &api.Response{
			Code: 304,
			Msg:  "create order fail",
		}, nil
	}
	return &api.Response{
		Code: 201,
		Msg:  "create order success",
	}, nil
}

var _ api.OrderServiceServer = &server{}

func (s *server) List(ctx context.Context, in *api.PageRequest) (out *api.ListResponse, err error) {
	/*buygoodsSpan, buygoodsCtx := opentracing.StartSpanFromContext(ctx, "BuyGoods")
	defer buygoodsSpan.Finish()
	if allow := s.Limit.Allow(); !allow {
		buygoodsSpan.SetTag("bgs.Limit.Allow()", false)
		//log.Println("limit cancel")
		return nil, errors.New("limit cancel")
	}
	buygoodsSpan.SetTag("bgs.Limit.Allow()", true)
	hystrixSpan, _ := opentracing.StartSpanFromContext(buygoodsCtx, "hystrix.Do")
	defer hystrixSpan.Finish()*/

	order := make([]*api.Order, 0)
	order = append(order, &api.Order{
		Id:                   1,
		Name:                 "1232123",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	order = append(order, &api.Order{
		Id:                   2,
		Name:                 "商品",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	order = append(order, &api.Order{
		Id:                   3,
		Name:                 "订单列表",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	out = &api.ListResponse{
		Order: order,
	}
	return
}
