package service

import (
	"google.golang.org/grpc"
	"orp/internal/app/myjob/conf"
	"orp/internal/app/myjob/dao"
)

type Service struct {
	c   *conf.TomlConfig
	dao *dao.Dao
	//autoid *autoid.AutoId
	autoIdConn *grpc.ClientConn
}

func New(c *conf.TomlConfig) *Service {
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
