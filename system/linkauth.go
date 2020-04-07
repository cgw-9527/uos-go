package system

import "github.com/tkblack/uos-go"

// NewLinkAuth creates an action from the `wxbio.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code uos.AccountName, actionName uos.ActionName, requiredPermission uos.PermissionName) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("linkauth"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      account,
				Permission: uos.PermissionName("active"),
			},
		},
		ActionData: uos.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     uos.AccountName    `json:"account"`
	Code        uos.AccountName    `json:"code"`
	Type        uos.ActionName     `json:"type"`
	Requirement uos.PermissionName `json:"requirement"`
}
