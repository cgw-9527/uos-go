package system

import (
	uos "github.com/uoscanada/uos-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner uos.AccountName) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("claimrewards"),
		Authorization: []uos.PermissionLevel{
			{Actor: owner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `uosio.system::claimrewards` action.
type ClaimRewards struct {
	Owner uos.AccountName `json:"owner"`
}
