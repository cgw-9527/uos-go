package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewRentCPU(
	from uos.AccountName,
	receiver uos.AccountName,
	loanPayment uos.Asset,
	loanFund uos.Asset,
) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("rentcpu"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(RentCPU{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentCPU struct {
	From        uos.AccountName
	Receiver    uos.AccountName
	LoanPayment uos.Asset
	LoanFund    uos.Asset
}
