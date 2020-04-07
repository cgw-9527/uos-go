package system

import (
	uos "github.com/tkblack/uos-go"
)

func NewBuyRAM(payer, receiver uos.AccountName, uosQuantity uint64) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("buyram"),
		Authorization: []uos.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: uos.NewUOSAsset(int64(uosQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `wxbio.system::buyram` action.
type BuyRAM struct {
	Payer    uos.AccountName `json:"payer"`
	Receiver uos.AccountName `json:"receiver"`
	Quantity uos.Asset       `json:"quant"` // specified in UOS
}
