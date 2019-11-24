package console

import (
	"github.com/qit-team/snow-core/command"
)

func RegisterCommand(c *command.Command) {
	c.AddFunc("test", test)
	c.AddFunc("insert_one_order", insertOne)
	c.AddFunc("insert_batch_order", InsertBatch)
	c.AddFunc("update_one_order", UpdateOne)
	c.AddFunc("update_batch_order", UpdateBatch)
	c.AddFunc("select_join", SelectJoin)
	c.AddFunc("use_session", UseSession)
	c.AddFunc("another_database", UseAnotherDatabase)
	c.AddFunc("show_sql_log", ShowSqlLog)
	c.AddFunc("use_sql", UseSql)
}
