package service

import (
	"bytes"
	"encoding/binary"
	"orp/internal/app/myshop/model"
	"strconv"
)

func (s *Service) GetShopOrderLog(tid string, refundId int64) model.TShopOrderLog {
	var order model.TShopOrderLog
	db := s.dao.Db
	if tid != "" {
		db.Where("tid = ?", tid)
	}
	if refundId != 0 {
		db.Where("refund_id = ?", refundId)
	}
	db.OrderBy("current_version desc").Get(&order)
	return order
}
func (s *Service) AddShopOrderLog(order model.TShopOrderLog) (int64, error) {

	id, err := s.GetID(s.c.App.ShopCode)
	if err != nil {
		return 0, err
	}
	order.Id = id
	order.CurrentVersion = 1
	if err != nil {
		return 0, err
	}
	var CurrentVersion uint8
	switch order.Type {
	case 1:
		maxCurrentVersion,err:=s.dao.Db.Query("SELECT MAX(`current_version`) as current_version FROM `t_shop_order_log` WHERE (tid = '"+order.Tid+"')")
		if err==nil {
			CurrentVersion = BytesToUint8(maxCurrentVersion[0]["current_version"])
		}
		//fmt.Println(d)
		//b, err = s.dao.Db.Where("tid = ?", order.Tid).Cols("current_version").OrderBy("current_version desc").Get(&getOrder)
	case 2:
		maxCurrentVersion,err:=s.dao.Db.Query("SELECT MAX(`current_version`) as current_version FROM `t_shop_order_log` WHERE (refund_id = '"+strconv.FormatInt(order.RefundId,10)+"')")
		if err==nil {
			CurrentVersion = BytesToUint8(maxCurrentVersion[0]["current_version"])
		}
		//b, err = s.dao.Db.Where("refund_id = ?", order.RefundId).Cols("current_version").OrderBy("current_version desc").Get(&getOrder)
	default:
	}
	order.CurrentVersion = CurrentVersion + 1
	i,err:=s.dao.Db.Insert(order)
	return i,err
}
func BytesToUint8(bys []byte) uint8 {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return uint8(data)
}
