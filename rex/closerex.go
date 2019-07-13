package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewCloseREX(owner uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("closerex"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(CloseREX{
			Ownwer: owner,
		}),
	}
}

type CloseREX struct {
	Ownwer uos.AccountName
}
