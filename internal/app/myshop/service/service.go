package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"orp/internal/app/myautoid_v2/api"
	"orp/internal/app/myshop/conf"
	"orp/internal/app/myshop/dao"
	"orp/pkg/consul"
	"time"
)

type Service struct {
	dao *dao.Dao
	autoIdConn *grpc.ClientConn
	c *conf.TomlConfig
}

func New(c *conf.TomlConfig) *Service {
	orderSnAutoIdConn,err:=consul.NewClientConn(c.App.CenterId, c.App.AutoIdService)
	if err != nil {
		log.Fatal(err)
	}
	return &Service{
		c:   c,
		dao: dao.New(c.Db),
		//autoid: autoid.New(c.Autoid),
		autoIdConn: orderSnAutoIdConn,
	}
}

func (s *Service) Close() {

}
func (s *Service) GetID(code int64) (id int64, err error) {
	client := api.NewAutoIdClient(s.autoIdConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(200)*time.Second)
	defer cancel()
	response, err := client.Get(ctx, &api.Request{
		Code: code,
	})
	return response.Id,err
}
