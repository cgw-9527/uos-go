package system

import "github.com/lialvin/uos-go"

// NewNonce returns a `nonce` action that lives on the
// `uosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `uos-bios` boot process by the
// `uosio.system` contract.
func NewNonce(nonce string) *uos.Action {
	a := &uos.Action{
		Account:       AN("uosio"),
		Name:          ActN("nonce"),
		Authorization: []uos.PermissionLevel{
			//{Actor: AN("uosio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}
