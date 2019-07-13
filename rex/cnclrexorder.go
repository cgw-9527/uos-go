package rex

import (
	uos "github.com/lialvin/uos-go"
)

func NewCancelREXOrder(owner uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("cnclrexorder"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(CancelREXOrder{
			Owner: owner,
		}),
	}
}

type CancelREXOrder struct {
	Owner uos.AccountName
}
