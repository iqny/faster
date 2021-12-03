package service

import (
	"orp/internal/app/myautoid_v2/conf"
	autoid "orp/pkg/autoid_v2"
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
