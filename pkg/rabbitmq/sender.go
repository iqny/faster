package rabbitmq

import (
	"bytes"
	"encoding/gob"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type Sender struct {
	conn      *amqp.Connection
	QueueName string //队列名称
	Job       string //任务名称
	Msg       []byte //消息体,回调时使用
	addr      string
	exchange  string
	mx        sync.Mutex
	queue     amqp.Queue
	ch        *amqp.Channel
}

func NweSender(addr, exchange string) *Sender {
	s := Sender{
		addr:     addr,
		exchange: exchange,
	}
	err := s.connect()
	if err != nil {
		log.Fatal(err)
	}
	return &s
}

func (s *Sender) Send(name, job string, sendMsg interface{}) (err error) {
	//连接断开了，进行重新连接
	s.mx.Lock()
	if s.conn == nil || s.conn.IsClosed() {
		err = s.connect()
		s.mx.Unlock()
		if err != nil {
			return err
		}
	} else {
		s.mx.Unlock()
	}
	s.QueueName = name
	s.Job = job
	s.Msg, err = byteEncoder(sendMsg)
	if err != nil {
		return err
	}
	if err = s.queueDeclare(); err != nil {
		return
	}
	var encResult bytes.Buffer
	enc := gob.NewEncoder(&encResult)
	if err = enc.Encode(s); err != nil {
		return err
	}

	err = s.ch.Publish(
		s.exchange,   // exchange
		s.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        encResult.Bytes(),
		})
	//log.Printf(" [x] Sent %s", enc_result.Bytes())
	return err
}

func (s *Sender) connect() (err error) {
	s.conn, err = amqp.Dial(s.addr)
	if err != nil {
		return
	}
	if err = s.channel();err!=nil{
		return
	}

	if err = s.exchangeDeclare();err!=nil{
		return
	}


	return
}
func (s *Sender) Close() {
	err := s.conn.Close()
	if err != nil {
		return 
	}
	s.ch.Close()
}
func byteEncoder(s interface{}) ([]byte, error) {
	var result bytes.Buffer
	enc := gob.NewEncoder(&result)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
func (s *Sender) queueDeclare() (err error) {
	s.queue, err = s.ch.QueueDeclare(
		s.QueueName, // name
		true,        // durable //是否持久化
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return
	}
	err = s.ch.QueueBind(
		s.queue.Name,
		//在pub/sub模式下，这里的key要为空
		s.queue.Name,
		s.exchange,
		false,
		nil,
	)
	return
}
func (s *Sender) exchangeDeclare() (err error) {
	err = s.ch.ExchangeDeclare(
		//交换机名称
		s.exchange,
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
	return
}
func (s *Sender) channel() (err error) {
	s.ch, err = s.conn.Channel()
	if err != nil {
		return err
	}
	return
}
