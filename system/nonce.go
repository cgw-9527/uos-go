package system

import "github.com/tkblack/uos-go"

// NewNonce returns a `nonce` action that lives on the
// `wxbio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `uos-bios` boot process by the
// `wxbio.system` contract.
func NewNonce(nonce string) *uos.Action {
	a := &uos.Action{
		Account:       AN("wxbio"),
		Name:          ActN("nonce"),
		Authorization: []uos.PermissionLevel{
			//{Actor: AN("wxbio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
