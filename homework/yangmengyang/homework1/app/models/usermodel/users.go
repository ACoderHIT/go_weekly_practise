package usermodel

import (
	"sync"
	"github.com/qit-team/snow-core/db"
)

var (
	once sync.Once
	m    *userModel
)
/**
 * Order实体，实体名必须和表名一致，因为底层获取的是结构体名称
 */
type Users struct {
	Id        int64     `xorm:"pk autoincr"` //注：使用getOne 或者ID() 需要设置主键
	UserName  string    `xorm:"varchar(20)"`
}

/**
 * 私有化，防止被外部new
 */
type userModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *userModel {
	once.Do(func() {
		m = new(userModel)
		//m.DiName = "" //设置数据库实例连接，默认db.SingletonMain
	})
	return m
}
