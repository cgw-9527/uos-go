package rex

import (
	uos "github.com/tkblack/uos-go"
)

func NewMoveToSavings(owner uos.AccountName, rex uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("mvtosavings"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(MoveToSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveToSavings struct {
	Owner uos.AccountName
	REX   uos.Asset
}
