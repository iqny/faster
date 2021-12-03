package autoid_v2

import (
	"fmt"
	"log"
	"math/rand"
	"orp/pkg/db"
	xtime "orp/pkg/time"
	"sync"
	"sync/atomic"
	"time"
	"xorm.io/xorm"
)

type Master struct {
	Dsn         string
	Active      int
	Idle        int
	AutoIdType  string
	AutoIdTable string
	IdleTimeout xtime.Duration
}
type Slave struct {
	Capacity int64
	Len      int64
}
type data struct {
	id       int64
	capacity int64
	num      int64
}

func (d *data) getId() int64 {
	//d.id++
	atomic.AddInt64(&d.id, +1)
	//d.num--
	atomic.AddInt64(&d.num, -1)
	if atomic.LoadInt64(&d.num) < 0 {
		return 0
	}
	return d.id
}
func (d *data) setId(id int64) {
	atomic.StoreInt64(&d.id, id)
	atomic.StoreInt64(&d.num, d.capacity)
}

type AutoId map[int][]*data

var mu sync.Mutex

func (a AutoId) GetAutoId(code int) (int64, error) {
	var id int64
	mu.Lock()
	defer mu.Unlock()
	//计算权重
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randWeight := r.Intn(int(conf.Slave.Len))
	if dat, ok := a[code]; ok {
		fmt.Println(dat)
		for i := 0; i < 2; i++ {
			obj:=dat[randWeight]
			id = obj.getId()
			fmt.Println("id",id)
			if id <= 0 {
				masterId, err := a.getMasterId(code, obj.capacity)
				fmt.Println(masterId)
				if err != nil {
					continue
				}
				obj.setId(masterId)
				continue
			} else {
				break
			}
		}
	}
	return id, nil
}
func (a AutoId) getMasterId(key int, count int64) (id int64, err error) {
	sql := fmt.Sprintf("update %s set id=last_insert_id(id+?) where k=?", conf.Db.AutoIdTable)
	res, err := engine.Exec(sql, count, key)
	if err != nil {
		//log.Fatal(err)
		return 0, err
	}
	//var id int64
	id, _ = res.LastInsertId()
	if id == 0 {
		sql := fmt.Sprintf("insert ignore into %s(k,id) values (?,0)", conf.Db.AutoIdTable)
		_, err := engine.Exec(sql, key)
		if err != nil {
			//log.Fatal(err)
			return 0, err
		}
		sql = fmt.Sprintf("update %s set id=last_insert_id(id+?) where k=?", conf.Db.AutoIdTable)
		res, err = engine.Exec(sql, count, key)
		if err != nil {
			//log.Fatal(err)
			return 0, err
		}
		id, _ = res.LastInsertId()
	}
	return
}

type Config struct {
	Db    *Master
	Slave *Slave
}

var engine *xorm.Engine
var conf *Config

func New(c *Config) *AutoId {

	dbConfig := &db.Config{
		DSN:         c.Db.Dsn,
		Drive:       "mysql",
		Active:      c.Db.Active,
		Idle:        c.Db.Idle,
		IdleTimeout: c.Db.IdleTimeout,
		Timezone:    "Asia/Shanghai",
	}

	conf = c
	engine = db.New(dbConfig)
	sql := fmt.Sprintf("select k from %s ", conf.Db.AutoIdTable)
	ks, err := engine.Query(sql)
	if err != nil {
		log.Fatalln(err)
	}
	autoId := AutoId{}

	for i := c.Slave.Len; i > 0; i-- {
		for kv := range ks {
			autoId[kv] = append(autoId[kv], &data{
				id:       0,
				capacity: c.Slave.Capacity,
				num:      0,
			})
		}
	}
fmt.Println(autoId[1005])
	return &autoId
}
func (a *AutoId) Close() {
	engine.Close()
}
