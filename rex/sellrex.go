package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewSellREX(from uos.AccountName, rex uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From uos.AccountName
	REX  uos.Asset
}
