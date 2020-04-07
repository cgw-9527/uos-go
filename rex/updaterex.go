package rex

import (
	uos "github.com/tkblack/uos-go"
)

func NewUpdateREX(owner uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner uos.AccountName
}
