package quenemodel

import (
	"github.com/qit-team/snow-core/db"
	"sync"
	"time"
)

var (
	once sync.Once
	m *queneLogModel
)

type QueneLog struct {
	Id           int `xorm:"pk autoincr"` //注：使用getOne 或者ID() 需要设置主键
	Data         string
	Update_time  time.Time
}

func (m *QueneLog) TableName() string {
	return "test_go"
}

type queneLogModel struct {
	db.Model
}

func GetInstance() *queneLogModel  {
	once.Do(func() {
		m = new(queneLogModel)
	})
	return m
}
