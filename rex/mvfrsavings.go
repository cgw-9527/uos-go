package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewMoveFromSavings(owner uos.AccountName, rex uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("mvfrsavings"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(MoveFromSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveFromSavings struct {
	Owner uos.AccountName
	REX   uos.Asset
}
