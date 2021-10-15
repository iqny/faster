package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"orp/pkg/rabbitmq"
	"orp/pkg/rabbitmq/logger"
	"path/filepath"
	"sync"
)

type TomlConfig struct {
	Queue  *rabbitmq.Config
	Logger *logger.Config //队列业务日志
}

//conf

var (
	confPath string
	Cfg      *TomlConfig
	once     sync.Once
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}
func Config() *TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs(confPath)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &Cfg); err != nil {
			panic(err)
		}
	})
	return Cfg
}

func watch()  {
	//监控到配置

}
