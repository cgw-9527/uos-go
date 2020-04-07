package rex

import (
	uos "github.com/tkblack/uos-go"
)

func NewRentNet(
	from uos.AccountName,
	receiver uos.AccountName,
	loanPayment uos.Asset,
	loanFund uos.Asset,
) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("rentnet"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(RentNet{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentNet struct {
	From        uos.AccountName
	Receiver    uos.AccountName
	LoanPayment uos.Asset
	LoanFund    uos.Asset
}
