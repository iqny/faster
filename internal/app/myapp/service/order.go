package service

import "orp/internal/app/myapp/model"

func (s *Service) Orders() []model.TOrder {
	var orders = make([]model.TOrder,0)
	s.dao.Db.OrderBy("order_id desc").Find(&orders)
	return orders
}
