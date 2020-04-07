package rex

import (
	uos "github.com/tkblack/uos-go"
)

func NewConsolidate(owner uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner uos.AccountName
}
