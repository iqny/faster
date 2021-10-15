package rabbitmq
type Config struct {
	Addr        string
	Work        []*WorkConfig
	LogFilePath string         //队列管理器日志
	Exchange    string         //交换机
}
type WorkConfig struct {
	Enable  bool
	Name    string
	MaxIdle int
}
