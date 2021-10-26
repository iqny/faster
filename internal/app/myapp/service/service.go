package service

import "orp/internal/app/myapp/dao"

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	return &Service{
		dao: dao.New("123"),
	}
}

func (s *Service) Close() {

}
