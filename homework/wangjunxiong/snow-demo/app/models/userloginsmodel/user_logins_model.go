package userloginsmodel

import (
	"github.com/qit-team/snow-core/db"
	"sync"
	"time"
)

var (
	once sync.Once
	m    *UserLoginsModel
)

//实体
type UserLogins struct {
	Device    string    `xorm:"'device' varchar(80)"`
	Id        int64     `xorm:"'id' bigint(20) pk autoincr"`
	Ip        string    `xorm:"'ip' varchar(15)"`
	LoginTime time.Time `xorm:"'login_time' timestamp"`
	Referer   string    `xorm:"'referer' varchar(128)"`
	UserId    int       `xorm:"'user_id' int(11)"`
	Useragent string    `xorm:"'useragent' varchar(255)"`
}

//表名
func (m *UserLogins) TableName() string {
	return "user_logins"
}

//私有化，防止被外部new
type UserLoginsModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *UserLoginsModel {
	once.Do(func() {
		m = new(UserLoginsModel)
		//m.DiName = "" //设置数据库实例连接，默认db.SingletonMain
	})
	return m
}

func (m *UserLoginsModel) InsertUserLogin(userLogin UserLogins) (id int64, err error) {
	id, err = m.Insert(userLogin);
	return
}

func (m *UserLoginsModel) InsertBatchUserLogin(userLogin []UserLogins) (id int64, err error) {
	id, err = m.Insert(userLogin);
	return
}

