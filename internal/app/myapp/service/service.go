package service

import (
	"orp/internal/app/myapp/conf"
	"orp/internal/app/myapp/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.TomlConfig) *Service {
	return &Service{
		dao: dao.New(c.Db),
	}
}

func (s *Service) Close() {
	s.dao.Db.Close()
}
