package system

import (
	uos "github.com/tkblack/uos-go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `wxbio.system` contract.
func NewDelegateBW(from, receiver uos.AccountName, stakeCPU, stakeNet uos.Asset, transfer bool) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("delegatebw"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: uos.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `wxbio.system::delegatebw` action.
type DelegateBW struct {
	From     uos.AccountName `json:"from"`
	Receiver uos.AccountName `json:"receiver"`
	StakeNet uos.Asset       `json:"stake_net"`
	StakeCPU uos.Asset       `json:"stake_cpu"`
	Transfer uos.Bool        `json:"transfer"`
}
