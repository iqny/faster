package db

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	xtime "orp/pkg/time"
	"os"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

type Config struct {
	Drive       string
	DSN         string // data source name.
	ShowSql     bool
	LogLevel    string
	LogFile     string
	Active      int            // pool
	Idle        int            // pool
	IdleTimeout xtime.Duration // connect max life time.
	Timezone    string
}

// New 创建一个db对象
func New(c *Config) (db *xorm.Engine) {
	db, err := xorm.NewEngine(c.Drive, c.DSN)
	if err != nil {
		log.Fatalf("db dsn(%s) error: %v", c.DSN, err)
		//panic(err)
	}
	location, err := time.LoadLocation(c.Timezone)
	db.TZLocation = location
	if err != nil {
		log.Printf("Fail set db Timezone %s", c.Timezone)
	}

	/*cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)//设置缓存
	Orm.SetDefaultCacher(cacher)*/

	db.SetMaxIdleConns(c.Idle)   //设置连接池的空闲数大小
	db.SetMaxOpenConns(c.Active) //设置最大打开连接数
	//fmt.Println(time.Duration(c.IdleTimeout))
	if c.IdleTimeout >0 {
		//db.SetConnMaxLifetime(time.Duration(c.IdleTimeout) / time.Second)
	}
	//打印到文件
	if c.LogFile != "" {
		logWriter, err := os.Create(c.LogFile)
		if err != nil {
			log.Printf("Fail to create xorm system logger: %v\n", err)
		}
		db.SetLogger(xlog.NewSimpleLogger(logWriter))
	}
	var level xlog.LogLevel
	switch c.LogLevel {
	case "debug":
		level = 0
	case "info":
		level = 1
	case "warning":
		level = 2
	case "err":
		level = 3
	}
	db.Logger().SetLevel(level)
	db.ShowSQL(c.ShowSql)
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}
