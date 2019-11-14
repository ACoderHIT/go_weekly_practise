package order

import (
	"time"
	"github.com/qit-team/snow-core/db"
	"sync"
)

var (
	once sync.Once
	m    *orderModel
)
/**
 * Order实体，实体名必须和表名一致，因为底层获取的是结构体名称
 */
type Orders struct {
	Id        int64     `xorm:"pk autoincr"` //注：使用getOne 或者ID() 需要设置主键
	UserId    int64
	OrderNo   string    `xorm:"varchar(20)"`
	Name      string    `xorm:"varchar(200)"`
	Status    int       `xorm:"int(11)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"` //此特性会激发软删除
}

/**
 * 表名规则
 * @wiki http://gobook.io/read/github.com/go-xorm/manual-zh-CN/chapter-02/3.tags.html
 */
//func (m *Orders) TableName() string {
//	return "orders"
//}

/**
 * 私有化，防止被外部new
 */
type orderModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *orderModel {
	once.Do(func() {
		m = new(orderModel)
		//m.DiName = "" //设置数据库实例连接，默认db.SingletonMain
	})
	return m
}

//func (m *bannerModel) insertO(pid int, limits ...int) (banners []*Banner, err error) {
//	banners = make([]*Banner, 0)
//	err = m.GetList(&banners, "pid = ?", []interface{}{pid}, limits)
//	return
//}