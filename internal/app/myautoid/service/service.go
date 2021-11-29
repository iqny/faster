package service

import (
	"orp/internal/app/myautoid/conf"
	"orp/pkg/autoid"
)

type Service struct {
	db *autoid.AutoId
}

func New(c *conf.TomlConfig) *Service {
	return &Service{
		db: autoid.New(c.Autoid),
	}
}

func (s *Service) Close() {
	s.db.Close()
}
