package system

import (
	uos "github.com/lialvin/uos-go"
)

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver uos.AccountName, bytes uint32) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("buyrambytes"),
		Authorization: []uos.PermissionLevel{
			{Actor: payer, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `uosio.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    uos.AccountName `json:"payer"`
	Receiver uos.AccountName `json:"receiver"`
	Bytes    uint32          `json:"bytes"`
}
