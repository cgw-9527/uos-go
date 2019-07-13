package system

import (
	uos "github.com/uoscanada/uos-go"
)

// NewRefund returns a `refund` action that lives on the
// `uosio.system` contract.
func NewRefund(owner uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("refund"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `uosio.system::refund` action
type Refund struct {
	Owner uos.AccountName `json:"owner"`
}
