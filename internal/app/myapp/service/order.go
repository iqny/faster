package service
//go:generate protoc -I. -I"%GOPATH%/src" -I"%GOPATH%/pkg/mod/github.com/gogo/protobuf@v1.3.2" --gofast_out=. "../api/order.proto"
import (
	"github.com/opentracing/opentracing-go"
	"golang.org/x/time/rate"
)

import (
	"context"
	"errors"
	"orp/internal/app/myapp/api"
)

type Order struct {
	Limit  *rate.Limiter
}


func (o *Order) List(ctx context.Context, in *api.PageRequest) (out *api.ListResponse, err error) {
	buygoodsSpan,buygoodsCtx:= opentracing.StartSpanFromContext(ctx,"BuyGoods")
	defer buygoodsSpan.Finish()
	if allow := o.Limit.Allow(); !allow {
		buygoodsSpan.SetTag("bgs.Limit.Allow()",false)
		//log.Println("limit cancel")
		return nil, errors.New("limit cancel")
	}
	buygoodsSpan.SetTag("bgs.Limit.Allow()",true)
	hystrixSpan,_:= opentracing.StartSpanFromContext(buygoodsCtx,"hystrix.Do")
	defer hystrixSpan.Finish()

	order := make([]*api.Order, 0)
	order = append(order,&api.Order{
		Id:                   1,
		Name:                 "1232123",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	order = append(order,&api.Order{
		Id:                   2,
		Name:                 "商品",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	order = append(order,&api.Order{
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
func (o Order) Add() {

}
