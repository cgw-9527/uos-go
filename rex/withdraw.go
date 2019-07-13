package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewWithdraw(owner uos.AccountName, amount uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("withdraw"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Withdraw{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Withdraw struct {
	Owner  uos.AccountName
	Amount uos.Asset
}
