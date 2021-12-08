package service

import (
	"context"
	"orp/internal/app/myautoid_v2/api"
	"orp/internal/app/myshop/conf"
	"orp/internal/app/myshop/dao"
	"orp/pkg/consul"
	"time"
)

type Service struct {
	dao *dao.Dao
	c *conf.TomlConfig
}

func New(c *conf.TomlConfig) *Service {
	return &Service{
		c:   c,
		dao: dao.New(c.Db),
	}
}

func (s *Service) Close() {

}
func (s *Service) GetID(code int64) (id int64, err error) {
	orderSnAutoIdConn,err:=consul.NewClientConn(s.c.App.CenterId, s.c.App.AutoIdService)
	if err != nil {
		return 0, err
	}
	client := api.NewAutoIdClient(orderSnAutoIdConn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	response, err := client.Get(ctx, &api.Request{
		Code: code,
	})
	return response.Id,err
}
