package system

import (
	"github.com/tkblack/uos-go"
)

func init() {
	uos.RegisterAction(AN("wxbio"), ActN("setcode"), SetCode{})
	uos.RegisterAction(AN("wxbio"), ActN("setabi"), SetABI{})
	uos.RegisterAction(AN("wxbio"), ActN("newaccount"), NewAccount{})
	uos.RegisterAction(AN("wxbio"), ActN("delegatebw"), DelegateBW{})
	uos.RegisterAction(AN("wxbio"), ActN("undelegatebw"), UndelegateBW{})
	uos.RegisterAction(AN("wxbio"), ActN("refund"), Refund{})
	uos.RegisterAction(AN("wxbio"), ActN("regproducer"), RegProducer{})
	uos.RegisterAction(AN("wxbio"), ActN("unregprod"), UnregProducer{})
	uos.RegisterAction(AN("wxbio"), ActN("regproxy"), RegProxy{})
	uos.RegisterAction(AN("wxbio"), ActN("voteproducer"), VoteProducer{})
	uos.RegisterAction(AN("wxbio"), ActN("claimrewards"), ClaimRewards{})
	uos.RegisterAction(AN("wxbio"), ActN("buyram"), BuyRAM{})
	uos.RegisterAction(AN("wxbio"), ActN("buyrambytes"), BuyRAMBytes{})
	uos.RegisterAction(AN("wxbio"), ActN("linkauth"), LinkAuth{})
	uos.RegisterAction(AN("wxbio"), ActN("unlinkauth"), UnlinkAuth{})
	uos.RegisterAction(AN("wxbio"), ActN("deleteauth"), DeleteAuth{})
	uos.RegisterAction(AN("wxbio"), ActN("rmvproducer"), RemoveProducer{})
	uos.RegisterAction(AN("wxbio"), ActN("setprods"), SetProds{})
	uos.RegisterAction(AN("wxbio"), ActN("setpriv"), SetPriv{})
	uos.RegisterAction(AN("wxbio"), ActN("canceldelay"), CancelDelay{})
	uos.RegisterAction(AN("wxbio"), ActN("bidname"), Bidname{})
	// uos.RegisterAction(AN("wxbio"), ActN("nonce"), &Nonce{})
	uos.RegisterAction(AN("wxbio"), ActN("sellram"), SellRAM{})
	uos.RegisterAction(AN("wxbio"), ActN("updateauth"), UpdateAuth{})
	uos.RegisterAction(AN("wxbio"), ActN("setramrate"), SetRAMRate{})
	uos.RegisterAction(AN("wxbio"), ActN("setalimits"), Setalimits{})
}

var AN = uos.AN
var PN = uos.PN
var ActN = uos.ActN
