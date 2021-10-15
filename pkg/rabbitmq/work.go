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

type Queue struct {
	conn     *amqp.Connection
	jobs     map[string]HandlerFunc
	c        *Config
	mx       sync.Mutex
	runCount map[string]int //运行时任务数
	logger   *logrus.Logger
	wg       sync.WaitGroup
	ctx      context.Context
}

func New(ctx context.Context, cfg *Config) *Queue {
	q := &Queue{c: cfg, ctx: ctx}
	err := q.connect()
	q.jobs = make(map[string]HandlerFunc)
	q.runCount = make(map[string]int)
	if err != nil {
		log.Fatal(err)
	}
	q.logger = newLog(cfg.LogFilePath)
	return q
}

// connect 连接rabbitmq
func (q *Queue) connect() (err error) {
	q.conn, err = amqp.Dial(q.c.Addr)
	if err != nil {
		return err
	}
	return nil
}

// Close 正常关闭
func (q *Queue) Close() {
	if q.conn != nil && !q.conn.IsClosed() {
		q.conn.Close()
	}
}

// Wait 等待其他goroutine退出
func (q *Queue) Wait() {
	q.wg.Wait()
}
func (q *Queue) Register(name string, call HandlerFunc) {
	q.jobs[name] = call
}

// Run 运行
func (q *Queue) Run() {
	go q.watch()
	for _, cfg := range q.c.Work {
		//fmt.Println(name.Name())
		if !cfg.Enable {
			continue
		}
		go q.master(cfg)
	}
}

// 监控重试连接
func (q *Queue) watch() {
	go func() {
		for {
			if q.conn == nil || q.conn.IsClosed() {
				q.mx.Lock()
				err := q.connect()
				q.logger.Warningln(fmt.Sprintf("retry connecting...:%v", err))
				q.mx.Unlock()
				if err != nil {
					time.Sleep(1 * time.Second)
				}
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

// 每个任务的父管理器
func (q *Queue) master(cfg *WorkConfig) {
	for {
		/*for i := cfg.MaxIdle; i > 0; i-- {
			go q.newShareQueue(cfg.Name)
		}*/
		q.mx.Lock()
		count := q.runCount[cfg.Name]
		q.mx.Unlock()
		//启动配置上的channel连接数
		if cfg.MaxIdle > count && (q.conn != nil && !q.conn.IsClosed()) {
			//fmt.Println("name:", cfg.Name, " count:", count)
			q.wg.Add(1)
			q.mx.Lock()
			q.runCount[cfg.Name]++
			q.mx.Unlock()
			go q.work(cfg.Name)
		} else {
			//fmt.Println(cfg.Name, cfg.MaxIdle, count, q.conn != nil, !q.conn.IsClosed())
			time.Sleep(1 * time.Second)

		}
	}
}

// work 连接渠道关联rabbitmq队列
func (q *Queue) work(queueName string) {

	defer func() {
		q.wg.Done()
		q.mx.Lock()
		q.runCount[queueName]--
		q.mx.Unlock()
	}()
	q.logger.Infoln(fmt.Sprintf("[*] Waiting for %s messages. To exit press CTRL+C", queueName))
	log.Printf("[*] Waiting for %s messages. To exit press CTRL+C", queueName)
	ch, err := q.conn.Channel()
	if err != nil {
		q.logger.Warningln(fmt.Sprintf("Failed to open a channel:%v", err))
		return
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		//交换机名称
		q.c.Exchange,
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
		q.logger.Warningln(fmt.Sprintf("Failed to ExchangeDeclare:%v", err))
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
		q.logger.Warningln(fmt.Sprintf("Failed to declare a queue:%v", err))
		return
	}
	err = ch.QueueBind(
		queue.Name,
		//在pub/sub模式下，这里的key要为空
		queue.Name,
		q.c.Exchange,
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
		q.logger.Warningln(fmt.Sprintf("Failed to register a consumer:%v", err))
		return
	}

	for {
		//协和任务退出
		if q.conn.IsClosed() {
			q.logger.Warningln("Failed connect to RabbitMQ")
			return
		}
		select {
		case msg := <-msgs:
			//接收消息
			var sender Sender
			decoder := gob.NewDecoder(bytes.NewReader(msg.Body))
			decoder.Decode(&sender)
			//查job处理方法
			if f, ok := q.jobs[sender.Job]; ok {
				v4 := strings.Split(uuid.NewV4().String(), "-")
				traceId := fmt.Sprintf("%s%s", v4[3], v4[4])
				q.logger.Infoln(fmt.Sprintf("%s Starting...", traceId))
				err := f(sender.Msg, traceId)
				q.logger.Infoln(fmt.Sprintf("%s Finished", traceId))
				if err == nil {
					err := msg.Ack(false)
					if err != nil {
						q.logger.Infoln(fmt.Sprintf("%s ack failed", traceId))
					} else {
						q.logger.Infoln(fmt.Sprintf("%s ack successes", traceId))
					}
				}
			}
		case <-q.ctx.Done():
			q.logger.Infoln(fmt.Sprintf("[*] Finished for %s messages.", queueName))
			//退出
			q.Close()
			return
		}
	}
}
