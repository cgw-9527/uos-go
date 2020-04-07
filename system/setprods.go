package system

import (
	uos "github.com/lialvin/uos-go"
	"github.com/lialvin/uos-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `wxbio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `uos-bios` boot process by the
// `wxbio.system` contract.
func NewSetProds(producers []ProducerKey) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("setprods"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("wxbio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `wxbio.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    uos.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}
