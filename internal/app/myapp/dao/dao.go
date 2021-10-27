package dao

import (
	"orp/pkg/db"
	"xorm.io/xorm"
)

type Dao struct {
	Db *xorm.Engine
}

func New(c *db.Config) *Dao {
	return &Dao{
		db.New(c),
	}
}
