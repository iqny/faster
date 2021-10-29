package main

import (
	"context"
	"log"
	"orp/internal/app/myapp/api"
	"orp/pkg/consul"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	//list()
	add()
}
func list() {
	conn, err := consul.NewClientConn("192.168.99.101:8500", "HelloService")
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

func add() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := consul.NewClientConn("192.168.99.101:8500", "HelloService")
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			c := api.NewOrderServiceClient(conn)
			for i := 0; i < 100000; i++ {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				//调用远端的方法
				_, err = c.Add(ctx, &api.Order{
					Id:   0,
					Name: strings.Join([]string{"vip"}, strconv.Itoa(i)),
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
}
