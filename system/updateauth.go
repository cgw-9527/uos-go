package system

import "github.com/tkblack/uos-go"

// NewUpdateAuth creates an action from the `wxbio.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account uos.AccountName, permission, parent uos.PermissionName, authority uos.Authority, usingPermission uos.PermissionName) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("updateauth"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: uos.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    uos.AccountName    `json:"account"`
	Permission uos.PermissionName `json:"permission"`
	Parent     uos.PermissionName `json:"parent"`
	Auth       uos.Authority      `json:"auth"`
}
