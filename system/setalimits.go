package system

import (
	uos "github.com/lialvin/uos-go"
)

// NewSetalimits sets the account limits. Requires signature from `uosio@active` account.
func NewSetalimits(account uos.AccountName, ramBytes, netWeight, cpuWeight int64) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("setalimit"),
		Authorization: []uos.PermissionLevel{
			{Actor: uos.AccountName("uosio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `uosio.system::setalimit` action.
type Setalimits struct {
	Account   uos.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}
