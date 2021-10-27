package autoid

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"orp/pkg/db"
	"strconv"
	"strings"
	"sync"
	"time"
	"xorm.io/xorm"
)

type Config struct {
	Master *Master
	Slave  []*Slave
}
type Master struct {
	Dsn         string
	Active      int
	Idle        int
	AutoIdType  string
	AutoIdTable string
	ShowSql     bool
	LogLevel    string
	LogFile     string
}
type Slave struct {
	Dsn         string
	Active      int
	Idle        int
	AutoIdType  string
	Capacity    []string
	Weight      int16
	MaxRetry    int16
	AutoIdTable string
	ShowSql     bool
	LogLevel    string
	LogFile     string
}

type AutoId struct {
	c    *Config
	lock sync.Mutex
	pool map[string]*xorm.Engine
}

var ormConfig *db.Config

func init() {
	ormConfig = &db.Config{
		DSN:         "",
		Drive:       "mysql",
		ShowSql:     false,
		LogLevel:    "",
		Active:      0,
		Idle:        0,
		IdleTimeout: 0,
		Timezone:    "Asia/Shanghai",
	}
}

var slaveDb *xorm.Engine
var masterDb *xorm.Engine

func New(c *Config) *AutoId {
	autoId := AutoId{c: c}
	autoId.pool = make(map[string]*xorm.Engine)
	for _, cnf := range c.Slave {
		if _, ok := autoId.pool[cnf.AutoIdTable]; ok {
			continue
		}
		ormConfig.Active = cnf.Active
		ormConfig.Idle = cnf.Idle
		ormConfig.DSN = cnf.Dsn
		//ormConfig.ShowSql = cnf.ShowSql
		//ormConfig.LogLevel = cnf.LogLevel
		//ormConfig.LogFile = cnf.LogFile
		autoId.pool[cnf.AutoIdTable] = db.New(ormConfig)
	}
	ormConfig.Active = c.Master.Active
	ormConfig.Idle = c.Master.Idle
	ormConfig.DSN = c.Master.Dsn
	//ormConfig.ShowSql = c.Master.ShowSql
	//ormConfig.LogLevel = c.Master.LogLevel
	//ormConfig.LogFile = c.Master.LogFile
	masterDb = db.New(ormConfig)
	autoId.pool[c.Master.AutoIdTable] = masterDb
	return &autoId
}
func (a *AutoId) GetAutoId(key int16) (id int64, err error) {
	var totalWeight int16
	for _, slave := range a.c.Slave {
		totalWeight += slave.Weight
	}
	length := int16(len(a.c.Slave) - 1)
	var count int16 = 1
	//var k int16
	//计算权重
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randWeight := r.Intn(int(totalWeight))
	var i int16 = 0
	j := int16(randWeight)
	for j >= 0 && i <= length {
		j = j - a.c.Slave[i].Weight
		i++
	}
	server := a.c.Slave[i]
	//fmt.Println("AutoIdTable", server.AutoIdTable, " i=", i)
	if server.AutoIdType == "master" {

	} else if server.AutoIdType == "slave" {
		id, err = a.getSlaveId(server, key, count)
		if id > 0 {
			return
		}
	} else if server.AutoIdType == "http" {

	}
	return
}

//master
func (a *AutoId) getMasterId(server *Master, key int16, count int64) int64 {

	sql := fmt.Sprintf("update %s set id=last_insert_id(id+?) where k=?", server.AutoIdTable)
	res, err := masterDb.Exec(sql, count, key)
	if err != nil {
		log.Fatal(err)
	}
	var id int64
	id, _ = res.LastInsertId()
	if id == 0 {
		sql := fmt.Sprintf("insert ignore into %s(k,id) values (?,0)", server.AutoIdTable)
		_, err := masterDb.Exec(sql, key)
		if err != nil {
			log.Fatal(err)
		}
		sql = fmt.Sprintf("update %s set id=last_insert_id(id+?) where k=?", server.AutoIdTable)
		res, err = masterDb.Exec(sql, count, key)
		if err != nil {
			log.Fatal(err)
		}
		id, _ = res.LastInsertId()
	}
	return id
}

func (a *AutoId) getSlaveId(server *Slave, key int16, count int16) (int64, error) {
	slaveDb = a.pool[server.AutoIdTable]
	max := server.MaxRetry
	for i := max; i > 0; i-- {
		id := a.slaveId(server, key, count)
		//fmt.Println("lastId=", id)
		if id > 0 {
			return id, nil
		} else if id == 0 {
			if i > 0 {
				var newCount int64 = 0
				var capacity = make(map[string]int64)
				for _, caps := range server.Capacity {
					nums := strings.Split(caps, ":")
					if len(nums) > 0 {
						num, err := strconv.Atoi(nums[1])
						if err != nil {
							num = 1000
						}
						capacity[nums[0]] = int64(num)
					}
				}
				var num int64
				var ok bool
				if num, ok = capacity[strconv.Itoa(int(key))]; ok {
					newCount = num
				} else {
					num, ok = capacity["default"]
					if !ok {
						log.Fatal("缺少default")
					}
					newCount = num
				}
				//parents := auto.c.Master
				a.lock.Lock()
				id := a.slaveId(server, key, count)
				if id > 0 {
					a.lock.Unlock()
					return id, nil
				}
				b, err := a.rechargeAutoIdSlave(server, a.c.Master, key, newCount, newCount)
				a.lock.Unlock()
				if b == true && err == nil {
					continue
				} else {
					//log.Fatal(err)
					return 0, err
				}
			}
		}
	}
	return 0, errors.New("找不到数据")
}
func (a *AutoId) rechargeAutoIdSlave(server *Slave, master *Master, key int16, capacity int64, minNum int64) (bool, error) {
	num, b := a.getAutoIdNum(server, key)
	if !b {
		return false, errors.New("slav找不到数据")
	}
	//fmt.Println("num==", num)
	if num < minNum || num == 0 {
		id := a.getMasterId(master, key, capacity)
		if id == 0 {
			return false, errors.New("master找不到数据")
		}
		sql := fmt.Sprintf("replace into %s(k,id,num) values('%d',%d-%d,%d)", server.AutoIdTable, key, id, capacity, capacity)
		db := a.pool[server.AutoIdTable]
		_, err := db.Exec(sql)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
func (a *AutoId) getAutoIdNum(server *Slave, key int16) (int64, bool) {
	db := a.pool[server.AutoIdTable]
	sql := fmt.Sprintf("select num from `%s` where k='%d' limit 1", server.AutoIdTable, key)
	res, err := db.Query(sql)
	if err != nil {
		//log.Fatal(err)
		return 0, false
	}
	if len(res) <= 0 {
		return 0, false
	}
	id, _ := strconv.Atoi(string(res[0]["num"]))
	return int64(id), true
}
func (a *AutoId) slaveId(server *Slave, key int16, count int16) int64 {
	sql := fmt.Sprintf("update %s set id=last_insert_id(id+?),num=num-? where k=? and num >= ?", server.AutoIdTable)
	//fmt.Println(sql)
	res, err := slaveDb.Exec(sql, count, count, key, count)

	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}
func (a *AutoId) Close() {
	for _, db := range a.pool {
		db.Close()
	}
}
