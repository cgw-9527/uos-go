package system

import (
	uos "github.com/uoscanada/uos-go"
)

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `uosio.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregprod`.
func NewRemoveProducer(producer uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("rmvproducer"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("uosio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `uosio.system::rmvproducer` action
type RemoveProducer struct {
	Producer uos.AccountName `json:"producer"`
}
