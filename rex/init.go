package rex

import uos "github.com/uoscanada/uos-go"

func init() {
	uos.RegisterAction(REXAN, ActN("buyrex"), BuyREX{})
	uos.RegisterAction(REXAN, ActN("closerex"), CloseREX{})
	uos.RegisterAction(REXAN, ActN("cnclrexorder"), CancelREXOrder{})
	uos.RegisterAction(REXAN, ActN("consolidate"), Consolidate{})
	uos.RegisterAction(REXAN, ActN("defcpuloan"), DefundCPULoan{})
	uos.RegisterAction(REXAN, ActN("defnetloan"), DefundNetLoan{})
	uos.RegisterAction(REXAN, ActN("deposit"), Deposit{})
	uos.RegisterAction(REXAN, ActN("fundcpuloan"), FundCPULoan{})
	uos.RegisterAction(REXAN, ActN("fundnetloan"), FundNetLoan{})
	uos.RegisterAction(REXAN, ActN("mvfrsavings"), MoveFromSavings{})
	uos.RegisterAction(REXAN, ActN("mvtosavings"), MoveToSavings{})
	uos.RegisterAction(REXAN, ActN("rentcpu"), RentCPU{})
	uos.RegisterAction(REXAN, ActN("rentnet"), RentNet{})
	uos.RegisterAction(REXAN, ActN("rexexec"), REXExec{})
	uos.RegisterAction(REXAN, ActN("sellrex"), SellREX{})
	uos.RegisterAction(REXAN, ActN("unstaketorex"), UnstakeToREX{})
	uos.RegisterAction(REXAN, ActN("updaterex"), UpdateREX{})
	uos.RegisterAction(REXAN, ActN("withdraw"), Withdraw{})
}

var AN = uos.AN
var PN = uos.PN
var ActN = uos.ActN

var REXAN = AN("uosio")
