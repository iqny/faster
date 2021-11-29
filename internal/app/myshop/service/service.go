package service

import (
	"orp/internal/app/myshop/conf"
	"orp/internal/app/myshop/dao"
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

}
