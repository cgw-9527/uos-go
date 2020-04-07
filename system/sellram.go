package system

import (
	uos "github.com/tkblack/uos-go"
)

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account uos.AccountName, bytes uint64) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("sellram"),
		Authorization: []uos.PermissionLevel{
			{Actor: account, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `wxbio.system::sellram` action.
type SellRAM struct {
	Account uos.AccountName `json:"account"`
	Bytes   uint64          `json:"bytes"`
}
