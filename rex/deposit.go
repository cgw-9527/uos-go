package rex

import (
	uos "github.com/tkblack/uos-go"
)

func NewDeposit(owner uos.AccountName, amount uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {
	Owner  uos.AccountName
	Amount uos.Asset
}
