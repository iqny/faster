package service

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"orp/internal/app/myapp/conf"
	"orp/internal/app/myapp/dao"
	"orp/internal/app/myautoid/api"
	"orp/pkg/consul"
	"time"
)

type Service struct {
	c   *conf.TomlConfig
	dao *dao.Dao
	//autoid *autoid.AutoId
	autoIdConn *grpc.ClientConn
}

func New(c *conf.TomlConfig) *Service {
	//defer conn.Close()
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
	s.dao.Db.Close()
	//s.autoid.Close()
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
