package main

import (
	"context"
	"log"
	"orp/internal/app/myapp/api"
	"orp/pkg/consul"
	"time"
)

func main() {
	conn, err := consul.NewClientConn("192.168.99.100:8500", "HelloService")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api.NewOrderServiceClient(conn)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		//调用远端的方法
		r, err := c.List(ctx, &api.PageRequest{
			Page: 1,
		})
		if err != nil {
			log.Printf("could not greet: %v", err)

		} else {
			//log.Printf("order: %s", r.Order)
			/*//获取的是系统时区的时间戳
			pbTimestamp := ptypes.TimestampNow()
			//此方法默认 UTC 时区
			goTime, _ := ptypes.Timestamp(pbTimestamp)
			//设定为系统时区
			fmt.Println(goTime.Local())*/
			for _, v := range r.Order {
				//设定为系统时区
				log.Printf("goods: %d,%s", v.Id, v.Name)
			}
		}
		time.Sleep(time.Second)
	}
}
