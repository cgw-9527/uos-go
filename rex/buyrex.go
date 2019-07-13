package rex

import (
	uos "github.com/lialvin/uos-go"
)

func NewBuyREX(from uos.AccountName, amount uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   uos.AccountName
	Amount uos.Asset
}
