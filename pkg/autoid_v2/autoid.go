package autoid_v2

import (
	"fmt"
	"log"
	"math/rand"
	"orp/pkg/db"
	xtime "orp/pkg/time"
	"strconv"
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randWeight := r.Intn(int(conf.Slave.Len))
	if dat, ok := a[code]; ok {
		for i := 0; i < 2; i++ {
			obj:=dat[randWeight]
			mu.Lock()
			id = obj.getId()
			if id <= 0 {
				masterId, err := a.getMasterId(code, obj.capacity)
				if err != nil {
					mu.Unlock()
					continue
				}
				obj.setId(masterId)
				mu.Unlock()
				continue
			} else {
				mu.Unlock()
				break
			}
		}
	}else{
		//请求不存在的code,自动生成并设置
		mu.Lock()
		defer mu.Unlock()
		if _, ok := a[code]; !ok {
			datas :=make([]*data,0)
			for i := conf.Slave.Len; i > 0; i-- {
				datas= append(datas, &data{
					id:       0,
					capacity: conf.Slave.Capacity,
					num:      0,
				})
			}
			a[code] = datas
		}
		dat = a[code]
		obj:=dat[randWeight]
		masterId, err := a.getMasterId(code, obj.capacity)
		if err != nil {
			return 0, err
		}
		obj.setId(masterId)
		id = obj.getId()
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
	var autoId = make(AutoId)
	for i := c.Slave.Len; i > 0; i-- {
		for _,code := range ks {
			ck,_:=strconv.Atoi(string(code["k"]))
			autoId[ck] = append(autoId[ck], &data{
				id:       0,
				capacity: c.Slave.Capacity,
				num:      0,
			})
		}
	}
	return &autoId
}
func (a *AutoId) Close() {
	engine.Close()
}
