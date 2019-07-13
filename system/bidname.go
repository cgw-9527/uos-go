package system

import (
	uos "github.com/lialvin/uos-go"
)

func NewBidname(bidder, newname uos.AccountName, bid uos.Asset) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("bidname"),
		Authorization: []uos.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `uosio.system_contract::bidname` action.
type Bidname struct {
	Bidder  uos.AccountName `json:"bidder"`
	Newname uos.AccountName `json:"newname"`
	Bid     uos.Asset       `json:"bid"` // specified in UOS
}
