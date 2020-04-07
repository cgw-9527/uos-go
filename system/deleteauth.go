package system

import "github.com/lialvin/uos-go"

// NewDeleteAuth creates an action from the `wxbio.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account uos.AccountName, permission uos.PermissionName) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("deleteauth"),
		Authorization: []uos.PermissionLevel{
			{Actor: account, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `wxbio.system` contract.
type DeleteAuth struct {
	Account    uos.AccountName    `json:"account"`
	Permission uos.PermissionName `json:"permission"`
}
