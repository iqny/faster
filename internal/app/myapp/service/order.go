package service

import (
	"errors"
	"orp/internal/app/myapp/model"
)

func (s *Service) Orders() []model.TOrder {
	var orders = make([]model.TOrder, 0)
	s.dao.Db.OrderBy("order_id desc").Find(&orders)
	return orders
}
func (s *Service) Add(name string) (int64, error) {
	id, err := s.autoid.GetAutoId(1005)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	order:=model.TOrder{
		OrderId: id,
		Name:    name,
	}
	return s.dao.Db.Insert(order)
}
