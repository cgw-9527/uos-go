package system

import uos "github.com/tkblack/uos-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `wxbio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `uos-bios` boot process by the
// `wxbio.system` contract.
func NewSetPriv(account uos.AccountName) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("setpriv"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("wxbio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(SetPriv{
			Account: account,
			IsPriv:  uos.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account uos.AccountName `json:"account"`
	IsPriv  uos.Bool        `json:"is_priv"`
}
