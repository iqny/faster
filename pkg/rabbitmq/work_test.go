package rabbitmq

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

var w *Queue

func TestNew(t *testing.T) {
	workConfig := []*WorkConfig{
		{
			Enable:  true,
			Name:    "testQueue",
			MaxIdle: 2,
		},
		{
			Enable:  true,
			Name:    "orderQueue",
			MaxIdle: 2,
		},
	}
	cfg := &Config{
		Addr:        "amqp://admin:admin@192.168.99.101:5672/",
		Work:        workConfig,
		LogFilePath: "D:/juest/src/orp/logs/queue/monitor.log",
		Exchange:    "test.direct",
	}

	ctx, cancel := context.WithCancel(context.Background())
	w = New(ctx,cfg)
	w.Register("testJob", func(args interface{}, taskId string) error {
		var scb string
		b := args.([]byte)
		decoder := gob.NewDecoder(bytes.NewReader(b))
		decoder.Decode(&scb)
		fmt.Println(scb, taskId)
		return nil
	})
	w.Register("orderJob", func(args interface{}, taskId string) error {
		var scb string
		b := args.([]byte)
		decoder := gob.NewDecoder(bytes.NewReader(b))
		decoder.Decode(&scb)
		fmt.Println(scb, taskId)
		return nil
	})
	w.Run()
	signalHandler(cancel)
}
func signalHandler(cancelFunc context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case sig := <-ch:
			fmt.Printf("get a signal %s, stop the apm-admin process\n", sig.String())
			switch sig {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				cancelFunc()
				w.Wait()
				fmt.Println("all obj work quit.")
				return
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}
}
func TestPush(t *testing.T) {
	client := NweSender("amqp://guest:guest@127.0.0.1:5672/", "test.direct")
	var data = make(map[string]interface{})
	data["no"] = 123
	for i := 0; i < 2; i++ {
		/*client.Send("testQueue", "testJob", "testJob...")
		client.Send("orderQueue", "orderJob", "orderJob...")
		client.Send("orderTransformQueue", "orderTransformJob", "orderTransformJob...")*/
		client.Send("pushWmsQueue", "pushWmsObj", data)
	}

}
