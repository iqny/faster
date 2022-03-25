package top_client

import (
	"fmt"
	"orp/pkg/sdk/topclient/request/tmallexchangereceivegetrequest"
	"testing"
)

func withAppKey(c *Config) {
	c.appKey = ""
}
func withAppSecret(c *Config) {
	c.appSecret = ""
}
func withFormat(c *Config)  {
	c.format = "json"
}
func withGatewayUrl(c*Config)  {
	c.gatewayUrl = "https://gw.api.taobao.com/router/rest"
}
func TestNew(t *testing.T) {
	cli := New(withAppKey, withAppSecret,withFormat,withGatewayUrl)
	req:= tmallexchangereceivegetrequest.New()
	req.SetFields("dispute_id, bizorder_id, num, buyer_nick, status, created, modified, reason, title, buyer_logistic_no, seller_logistic_no, bought_sku, exchange_sku, buyer_address, address, buyer_phone, buyer_logistic_name, seller_logistic_name, alipay_no, buyer_name, seller_nick,time_out,good_status,payment,price,biz_order_id,desc")
	req.SetPageSize("50")
	req.SetPageNo("1")
	req.SetStartGmtModifiedTime("2022-03-23 00:00:00")
	req.SetEndGmtModifiedTime("2022-03-24 14:00:00")
	res, err := cli.Execute(req,"")
	if err == nil {
		fmt.Println(res)
	}
}

