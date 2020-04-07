package system

import (
	uos "github.com/tkblack/uos-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner uos.AccountName) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
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

// ClaimRewards represents the `wxbio.system::claimrewards` action.
type ClaimRewards struct {
	Owner uos.AccountName `json:"owner"`
}
