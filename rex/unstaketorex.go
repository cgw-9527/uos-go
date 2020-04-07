package rex

import (
	uos "github.com/tkblack/uos-go"
)

func NewUnstakeToREX(
	owner uos.AccountName,
	receiver uos.AccountName,
	fromNet uos.Asset,
	fromCPU uos.Asset,
) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("unstaketorex"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(UnstakeToREX{
			Owner:    owner,
			Receiver: receiver,
			FromNet:  fromNet,
			FromCPU:  fromCPU,
		}),
	}
}

type UnstakeToREX struct {
	Owner    uos.AccountName
	Receiver uos.AccountName
	FromNet  uos.Asset
	FromCPU  uos.Asset
}
