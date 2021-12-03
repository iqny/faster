package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	autoid "orp/pkg/autoid_v2"
	"path/filepath"
	"sync"
)

type TomlConfig struct {
	App *AppConfig
	Autoid   *autoid.Config
}
type AppConfig struct {
	ServiceName string
	Port int
	CenterId string
}

//conf

var (
	confPath string
	//Cfg      *TomlConfig
	once     sync.Once
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}
func Config() (c *TomlConfig) {
	once.Do(func() {
		filePath, err := filepath.Abs(confPath)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &c); err != nil {
			//panic(err)
		}
	})
	return
}
