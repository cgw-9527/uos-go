package system

import "github.com/lialvin/uos-go"

// NewUnlinkAuth creates an action from the `wxbio.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code uos.AccountName, actionName uos.ActionName) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("unlinkauth"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      account,
				Permission: uos.PermissionName("active"),
			},
		},
		ActionData: uos.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account uos.AccountName `json:"account"`
	Code    uos.AccountName `json:"code"`
	Type    uos.ActionName  `json:"type"`
}
