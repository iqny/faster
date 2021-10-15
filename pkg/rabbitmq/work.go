package rabbitmq

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"log"
	"strings"
	"sync"
	"time"
)

type HandlerFunc func(args interface{}, taskId string) error

func (h HandlerFunc) Call(args interface{}, taskId string) error {
	return h(args, taskId)
}

type Work struct {
	conn     *amqp.Connection
	jobs     map[string]HandlerFunc
	c        *Config
	mx       sync.Mutex
	runCount map[string]int //运行时任务数
	logger   *logrus.Logger
	wg       sync.WaitGroup
	ctx      context.Context
}

func New(ctx context.Context, cfg *Config) *Work {
	w := &Work{c: cfg, ctx: ctx}
	err := w.connect()
	w.jobs = make(map[string]HandlerFunc)
	w.runCount = make(map[string]int)
	if err != nil {
		log.Fatal(err)
	}
	w.logger = newLog(cfg.LogFilePath)
	return w
}

// connect 连接rabbitmq
func (w *Work) connect() (err error) {
	w.conn, err = amqp.Dial(w.c.Addr)
	if err != nil {
		return err
	}
	return nil
}
func (w *Work) Close() {
	if w.conn != nil && !w.conn.IsClosed() {
		w.conn.Close()
	}

}
func (w *Work) Wait() {
	w.wg.Wait()
}
func (w *Work) Register(name string, call HandlerFunc) {
	w.jobs[name] = call
}
func (w *Work) Run() {
	go w.watch()
	for _, cfg := range w.c.Work {
		//fmt.Println(name.Name())
		if !cfg.Enable {
			continue
		}
		go w.master(cfg)
	}
}

// 监控重试连接
func (w *Work) watch() {
	go func() {
		for {
			if w.conn == nil || w.conn.IsClosed() {
				w.mx.Lock()
				err := w.connect()
				w.logger.Warningln(fmt.Sprintf("retry connecting...:%v", err))
				w.mx.Unlock()
				if err != nil {
					time.Sleep(1 * time.Second)
				}
			}else{
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (w *Work) master(cfg *WorkConfig) {
	for {
		/*for i := cfg.MaxIdle; i > 0; i-- {
			go w.newShareQueue(cfg.Name)
		}*/
		w.mx.Lock()
		count := w.runCount[cfg.Name]
		w.mx.Unlock()
		if cfg.MaxIdle > count && (w.conn != nil && !w.conn.IsClosed()) {
			//fmt.Println("name:", cfg.Name, " count:", count)
			w.wg.Add(1)
			w.mx.Lock()
			w.runCount[cfg.Name]++
			w.mx.Unlock()
			go w.newShareQueue(cfg.Name)
		} else {
			//fmt.Println(cfg.Name, cfg.MaxIdle, count, w.conn != nil, !w.conn.IsClosed())
			time.Sleep(1 * time.Second)

		}
	}
}
func (w *Work) newShareQueue(queueName string) {

	defer func() {
		w.wg.Done()
		w.mx.Lock()
		w.runCount[queueName]--
		w.mx.Unlock()
	}()
	w.logger.Infoln(fmt.Sprintf("[*] Waiting for %s messages. To exit press CTRL+C", queueName))
	log.Printf("[*] Waiting for %s messages. To exit press CTRL+C", queueName)
	ch, err := w.conn.Channel()
	if err != nil {
		w.logger.Warningln(fmt.Sprintf("Failed to open a channel:%v", err))
		return
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		//交换机名称
		w.c.Exchange,
		//交换机类型 广播类型
		"direct",
		//是否持久化
		true,
		//是否字段删除
		false,
		//true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		//是否阻塞 true表示要等待服务器的响应
		false,
		nil,
	)
	if err != nil {
		w.logger.Warningln(fmt.Sprintf("Failed to ExchangeDeclare:%v", err))
	}
	//checkOnError(err, "Failed to open a channel")
	queue, err := ch.QueueDeclare(
		queueName,
		true, //是否持久化
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		w.logger.Warningln(fmt.Sprintf("Failed to declare a queue:%v", err))
		return
	}
	err = ch.QueueBind(
		queue.Name,
		//在pub/sub模式下，这里的key要为空
		queue.Name,
		w.c.Exchange,
		false,
		nil,
	)
	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		w.logger.Warningln(fmt.Sprintf("Failed to register a consumer:%v", err))
		return
	}

	for {
		//协和任务退出
		if w.conn.IsClosed() {
			w.logger.Warningln("Failed connect to RabbitMQ")
			return
		}
		select {
		case msg := <-msgs:
			//接收消息
			var sender Sender
			decoder := gob.NewDecoder(bytes.NewReader(msg.Body))
			decoder.Decode(&sender)
			//查job处理方法
			if f, ok := w.jobs[sender.Job]; ok {
				v4 := strings.Split(uuid.NewV4().String(), "-")
				traceId := fmt.Sprintf("%s%s", v4[3], v4[4])
				w.logger.Infoln(fmt.Sprintf("%s Starting...", traceId))
				err := f(sender.Msg, traceId)
				w.logger.Infoln(fmt.Sprintf("%s Finished", traceId))
				if err == nil {
					err := msg.Ack(false)
					if err != nil {
						w.logger.Infoln(fmt.Sprintf("%s ack failed", traceId))
					} else {
						w.logger.Infoln(fmt.Sprintf("%s ack successes", traceId))
					}
				}
			}
		case <-w.ctx.Done():
			w.logger.Infoln(fmt.Sprintf("[*] Finished for %s messages.", queueName))
			//退出
			w.Close()
			return
		}
	}
}
