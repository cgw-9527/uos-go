package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewREXExec(user uos.AccountName, max uint16) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []uos.PermissionLevel{
			{Actor: user, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User uos.AccountName
	Max  uint16
}
