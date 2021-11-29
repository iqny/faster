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
	orderSn, err :=s.GetID(s.c.App.OrderCode)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}

	order := model.TOrder{
		OrderId: orderSn,
		Name:    name,
	}
	return s.dao.Db.Insert(order)
}
/*func (s *Service) Add(name string) (int64, error) {
	orderSn, err := s.autoid.GetAutoId(1005)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	uid, err := s.autoid.GetAutoId(1006)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	tid, err := s.autoid.GetAutoId(1007)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	oid1, err := s.autoid.GetAutoId(1009)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	oid2, err := s.autoid.GetAutoId(1009)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	id, err := s.autoid.GetAutoId(1008)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	id2, err := s.autoid.GetAutoId(1008)
	if err != nil {
		return 0, errors.New("添加订单失败")
	}
	oidStr1 := strconv.FormatInt(oid1, 10)
	oidStr2 := strconv.FormatInt(oid2, 10)
	tidStr := strconv.FormatInt(tid, 10)
	oi1:=rand.Intn(3)
	g:=rand.Intn(3)
	orders := []model.TOrder{
		{
			OrderSn: orderSn,
			UserId:    uid,
			Tid:       tidStr,
			Consignee: "张三",
			Country:   "中国",
			Province:  "广东省",
			District:  "广州市",
			Town:      "天河区",
			Address:   "体育中心1688号",
			Tel:       "137612312313",
		},
		{
			OrderSn: orderSn,
			UserId:    uid,
			Tid:       tidStr,
			Consignee: "李四",
			Country:   "中国",
			Province:  "广东省",
			District:  "广州市",
			Town:      "天河区",
			Address:   "体育中心1688号",
			Tel:       "13658746857",
		},
		{
			OrderSn: orderSn,
			UserId:    uid,
			Tid:       tidStr,
			Consignee: "麻子",
			Country:   "中国",
			Province:  "广东省",
			District:  "广州市",
			Town:      "天河区",
			Address:   "体育中心1688号",
			Tel:       "1362577895",
		},
		{
			OrderSn: orderSn,
			UserId:    uid,
			Tid:       tidStr,
			Consignee: "页子烟气",
			Country:   "中国",
			Province:  "广东省",
			District:  "广州市",
			Town:      "天河区",
			Address:   "体育中心1688号",
			Tel:       "13985745624",
		},
	}
	/*order := model.TOrder{
		OrderSn:   orderSn,
		UserId:    uid,
		Tid:       tidStr,
		Consignee: "洋气",
		Country:   "中国",
		Province:  "广东省",
		District:  "广州市",
		Town:      "天河区",
		Address:   "体育中心1688号",
		Tel:       "137612312313231",
	}
	goods1 := []model.TOrderItem{
		{
			Id: id2,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "天猫旗舰店销售破2亿,1分18秒完成首单,谁在看好1919?",
			GoodsNumber:      "1",
			MarketPrice:      1000.89,
			Discount:         100.5,
			TransactionPrice: 10.5,
			Oid:              oidStr1,
		},
		{
			Id: id2,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "天猫精选-理想生活上天猫",
			GoodsNumber:      "2",
			MarketPrice:      300.89,
			Discount:         10.9,
			TransactionPrice: 101.5,
			Oid:              oidStr1,
		},
		{
			Id: id2,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "买天猫-旗舰店上万能的淘宝!优享品质,惊喜价格!",
			GoodsNumber:      "3",
			MarketPrice:      1001.89,
			Discount:         101.5,
			TransactionPrice: 20.5,
			Oid:              oidStr1,
		},
		{
			Id: id2,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "天猫的最新相关信息",
			GoodsNumber:      "1",
			MarketPrice:      100.89,
			Discount:         10.5,
			TransactionPrice: 0.5,
			Oid:              oidStr1,
		},
	}
	goods2 := []model.TOrderItem{
		{
			Id: id,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "我的天猫店铺需要提升基础考核各指标，应该怎么...",
			GoodsNumber:      "1",
			MarketPrice:      100.89,
			Discount:         10.5,
			TransactionPrice: 0.5,
			Oid:              oidStr2,
		},
		{
			Id: id,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "Go语言Mock使用基本指南详解_Golang_脚本之家",
			GoodsNumber:      "1",
			MarketPrice:      100.89,
			Discount:         10.5,
			TransactionPrice: 0.5,
			Oid:              oidStr2,
		},
		{
			Id: id,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "Go Mock (gomock)简明教程 | 快速入门 | 极客兔兔",
			GoodsNumber:      "1",
			MarketPrice:      100.89,
			Discount:         10.5,
			TransactionPrice: 0.5,
			Oid:              oidStr2,
		},
		{
			Id: id,
			OrderSn:          orderSn,
			UserId:           uid,
			Tid:              tidStr,
			GoodsName:        "golang三大基础mock大法 - ExplorerMan - 博客园",
			GoodsNumber:      "1",
			MarketPrice:      100.89,
			Discount:         10.5,
			TransactionPrice: 0.5,
			Oid:              oidStr2,
		},
	}
	goods:= []model.TOrderItem{
		goods1[g],
		goods2[g],
	}
	s.dao.Db.Insert(orders[oi1])
	return s.dao.Db.Insert(goods)
}*/
