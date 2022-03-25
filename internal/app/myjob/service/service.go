package service

import (
	"google.golang.org/grpc"
	"orp/internal/app/myjob/conf"
	"orp/internal/app/myjob/dao"
	"orp/pkg/jaeger"
)

type Service struct {
	c   *conf.TomlConfig
	dao *dao.Dao
	//autoid *autoid.AutoId
	autoIdConn *grpc.ClientConn
}

func New(c *conf.TomlConfig) *Service {
	_, _, err := jaeger.NewTracer("test_v1", "127.0.0.1:6831")
	if err != nil {
		return nil
	}
	//defer i.Close()
	return &Service{
		c:   c,
		dao: dao.New(c.Db),
		//autoid: autoid.New(c.Autoid),
	}
}

func (s *Service) Close() {
	s.dao.Db.Close()
	//s.autoid.Close()
}
