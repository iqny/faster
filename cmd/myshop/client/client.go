package main

import (
	"context"
	"fmt"
	"log"
	api2 "orp/internal/app/myautoid_v2/api"
	"orp/internal/app/myshop/api"
	"orp/pkg/consul"
	"strconv"
	"sync"
	"time"
)

func main() {
	//list()
	add()
}


func add() {
	var wg sync.WaitGroup
	conn, err := consul.NewClientConn("172.16.51.95:8500", "ShopService")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api.NewShopClient(conn)
	connAuto, err := consul.NewClientConn("172.16.51.95:8500", "AutoIdV2Service")
	if err != nil {
		log.Fatal(err)
	}
	defer connAuto.Close()
	c1 := api2.NewAutoIdClient(connAuto)

s:=time.Now()
	//调用远端的方法
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				ctx, cancel := context.WithTimeout(context.Background(), time.Duration(200)*time.Second)
				ctx1, cancel1 := context.WithTimeout(context.Background(), time.Duration(200)*time.Second)
				defer cancel()
				defer cancel1()
				res,_:=c1.Get(ctx1,&api2.Request{
					Code:                 1010,
				})
				_, err = c.AddShopOrderLog(ctx, &api.Order{
					ShopId:               1,
					Tid:                  strconv.FormatInt(res.GetId(),10),
					RefundId:             0,
					Type:                 1,
					Content:              "{\"abc\":\"kir\"}",
				})
				if err != nil {
					log.Printf("could not greet: %v", err)
				} else {
					//fmt.Println(r.Code)
				}
			}
		}()

	}
	wg.Wait()
	n:=time.Now()
	fmt.Println(n.Sub(s))
}
