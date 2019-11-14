package console

import (
	"github.com/qit-team/snow-core/command"
)

func RegisterCommand(c *command.Command) {
	c.AddFunc("test", test)
	c.AddFunc("yangtest", yangtest)
}
