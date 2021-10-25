package wms

import (
	"orp/pkg/wms/interface_factory"
	taotao "orp/pkg/wms/taotao/client"
)

func New(wmsName string, cfs ...interface_factory.ConfigFunc) interface_factory.Wms {
	switch wmsName {
	case "taobao":
		c := taotao.New(cfs...)
		return c
	}
	return nil
}
