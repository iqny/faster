package request

import "orp/pkg/taobaosdk/wms/response"

type Request interface {
	ToXML() string
	Check() (response.Response,error)
	GetMethod() string
}
