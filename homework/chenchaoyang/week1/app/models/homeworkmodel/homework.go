package homeworkmodel

import (
	"github.com/qit-team/snow-core/db"
	"sync"
	"time"
)

/**
 * Banner实体
 */
type Orders struct {
	Id        int64 `xorm:"pk autoincr"` //注：使用getOne 或者ID() 需要设置主键
	OrderNo   string
	UserId    string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

var once sync.Once
var singleIns *homeworkModel

/**
 * 私有化，防止被外部new
 */
type homeworkModel struct {
	db.Model //组合基础Model，集成基础Model的属性和方法
}

//单例模式
func GetInstance() *homeworkModel {
	once.Do(func() {
		singleIns = new(homeworkModel)
	})
	return singleIns
}
