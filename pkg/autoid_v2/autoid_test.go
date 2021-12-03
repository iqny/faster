package autoid_v2

import (
	"flag"
	"fmt"
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
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1; j++ {
				id, err := auto.GetAutoId(1005)
				fmt.Println(id,"===",err)
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
	fmt.Println(confPath)
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
