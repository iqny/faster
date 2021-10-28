package autoid

import (
	"flag"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
	"testing"
)

func TestNewAutoID(t *testing.T) {
	flag.Parse()
	c := Config1()
	auto := New(c.Autoid)
	//id,err:=auto.GetAutoId(1005)
	//fmt.Println("id:",id,"err:",err)
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 20; j++ {
				auto.GetAutoId(1005)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	auto.Close()
}
func BenchmarkAutoId_GetAutoId(b *testing.B) {
	b.ResetTimer()
	flag.Parse()
	c:=Config1()
	auto:=New(c.Autoid)
	for i:=0;i<b.N;i++{
		auto.GetAutoId(1005)
	}
}
type TomlConfig struct {
	Autoid   *Config
}
var (
	confPath string
	Cfg      *TomlConfig
	once     sync.Once
)

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}
func Config1() *TomlConfig {
	once.Do(func() {
		filePath, err := filepath.Abs(confPath)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &Cfg); err != nil {
			//panic(err)
		}
	})
	return Cfg
}
