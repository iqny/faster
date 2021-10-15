package consul

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"orp/pkg/consul/register"
	"orp/pkg/consul/resolver"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

//NewClientConn 创建客户端
func NewClientConn(consulAddr, serviceName string) (conn *grpc.ClientConn, err error) {
	schema, err := resolver.GenerateAndRegisterConsulResolver(consulAddr, serviceName)
	if err != nil {
		err = errors.New(fmt.Sprintf("init consul resovler err:%v", err))
	}
	//建立连接
	conn, err = grpc.Dial(fmt.Sprintf("%s:///%s", schema, serviceName), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		err = errors.New(fmt.Sprintf("did not connect: %v", err))
	}
	return
}

//Registry 服务注册自定义结构体
type Registry struct {
	consulAddr, service string
	port                int
	listener            net.Listener
	Server              *grpc.Server
	register *register.ConsulRegister
}

//NewRegister 创建新的服务注册
func NewRegister(consulAddr, service string, port int) (*Registry, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil,err
	}
	addrs := strings.Split(listener.Addr().String(), ":")
	port, err = strconv.Atoi(addrs[len(addrs)-1])
	if err != nil {
		return nil,err
	}
	log.Println("start server port :", addrs[len(addrs)-1])
	//consul service register
	nr := register.NewConsulRegister(consulAddr, service, port)
	nr.Register()
	//start grpc server
	serv := grpc.NewServer()
	//registe health check
	grpc_health_v1.RegisterHealthServer(serv, &register.HealthImpl{})

	return &Registry{consulAddr: consulAddr, service: service, port: port, listener: listener, Server: serv, register: nr}, nil
}
//Run 启动
func (r *Registry)Run()  {
	//server hook
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
		log.Println("do run hook")
		r.register.Deregister()
		r.Server.Stop()
	}()

	if err := r.Server.Serve(r.listener); err != nil {
		panic(err)
	}
}
