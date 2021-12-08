package main

import (
	"context"
	"fmt"
	"log"
	"orp/internal/app/myautoid/api"
	"orp/pkg/consul"
	"sync"
	"time"
)

func main() {
	//list()
	s:=time.Now()
	add()
	fmt.Println(time.Now().Sub(s))

}

func add() {
	var wg sync.WaitGroup
	conn, err := consul.NewClientConn("127.0.0.1:8500", "AutoIdV2Service")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api.NewAutoIdClient(conn)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(200)*time.Second)
	//调用远端的方法
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				_, err = c.Get(ctx, &api.Request{
					Code:                 1011,
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
