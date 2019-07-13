package system

import (
	uos "github.com/lialvin/uos-go"
	"github.com/lialvin/uos-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `uosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `uos-bios` boot process by the
// `uosio.system` contract.
func NewSetProds(producers []ProducerKey) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("setprods"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("uosio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `uosio.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    uos.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}
