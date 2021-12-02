package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"orp/pkg/autoid"
	"orp/pkg/db"
	"path/filepath"
	"sync"
)

type TomlConfig struct {
	App *AppConfig
	Db  *db.Config
	Autoid *autoid.Config
}
type AppConfig struct {
	Timezone      string
	Host          string
	Gzip          bool
	AutoIdService string
	ShopCode     int64
	ServiceName string
	Port int
	CenterId string
}

//conf

var (
	confPath string
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
