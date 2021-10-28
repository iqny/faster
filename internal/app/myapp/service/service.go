package service

import (
	"orp/internal/app/myapp/conf"
	"orp/internal/app/myapp/dao"
	"orp/pkg/autoid"
)

type Service struct {
	dao *dao.Dao
	autoid *autoid.AutoId
}

func New(c *conf.TomlConfig) *Service {
	return &Service{
		dao: dao.New(c.Db),
		autoid: autoid.New(c.Autoid),
	}
}

func (s *Service) Close() {
	s.dao.Db.Close()
	s.autoid.Close()
}
