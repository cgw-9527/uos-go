package system

import (
	"github.com/uoscanada/uos-go"
)

func init() {
	uos.RegisterAction(AN("uosio"), ActN("setcode"), SetCode{})
	uos.RegisterAction(AN("uosio"), ActN("setabi"), SetABI{})
	uos.RegisterAction(AN("uosio"), ActN("newaccount"), NewAccount{})
	uos.RegisterAction(AN("uosio"), ActN("delegatebw"), DelegateBW{})
	uos.RegisterAction(AN("uosio"), ActN("undelegatebw"), UndelegateBW{})
	uos.RegisterAction(AN("uosio"), ActN("refund"), Refund{})
	uos.RegisterAction(AN("uosio"), ActN("regproducer"), RegProducer{})
	uos.RegisterAction(AN("uosio"), ActN("unregprod"), UnregProducer{})
	uos.RegisterAction(AN("uosio"), ActN("regproxy"), RegProxy{})
	uos.RegisterAction(AN("uosio"), ActN("voteproducer"), VoteProducer{})
	uos.RegisterAction(AN("uosio"), ActN("claimrewards"), ClaimRewards{})
	uos.RegisterAction(AN("uosio"), ActN("buyram"), BuyRAM{})
	uos.RegisterAction(AN("uosio"), ActN("buyrambytes"), BuyRAMBytes{})
	uos.RegisterAction(AN("uosio"), ActN("linkauth"), LinkAuth{})
	uos.RegisterAction(AN("uosio"), ActN("unlinkauth"), UnlinkAuth{})
	uos.RegisterAction(AN("uosio"), ActN("deleteauth"), DeleteAuth{})
	uos.RegisterAction(AN("uosio"), ActN("rmvproducer"), RemoveProducer{})
	uos.RegisterAction(AN("uosio"), ActN("setprods"), SetProds{})
	uos.RegisterAction(AN("uosio"), ActN("setpriv"), SetPriv{})
	uos.RegisterAction(AN("uosio"), ActN("canceldelay"), CancelDelay{})
	uos.RegisterAction(AN("uosio"), ActN("bidname"), Bidname{})
	// uos.RegisterAction(AN("uosio"), ActN("nonce"), &Nonce{})
	uos.RegisterAction(AN("uosio"), ActN("sellram"), SellRAM{})
	uos.RegisterAction(AN("uosio"), ActN("updateauth"), UpdateAuth{})
	uos.RegisterAction(AN("uosio"), ActN("setramrate"), SetRAMRate{})
	uos.RegisterAction(AN("uosio"), ActN("setalimits"), Setalimits{})
}

var AN = uos.AN
var PN = uos.PN
var ActN = uos.ActN
