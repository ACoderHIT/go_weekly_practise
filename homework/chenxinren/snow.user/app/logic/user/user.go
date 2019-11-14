package user

import (
	"context"
	"snow.user/app/logic/user/processes"
)

func DonEntry(c context.Context, msg string) (ok bool, err error) {
	return processes.Handle(c, msg)
}