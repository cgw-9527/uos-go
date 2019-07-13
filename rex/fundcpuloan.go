package rex

import (
	uos "github.com/lialvin/uos-go"
)

func NewFundCPULoan(from uos.AccountName, loanNumber uint64, payment uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("fundcpuloan"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(FundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundCPULoan struct {
	From       uos.AccountName
	LoanNumber uint64
	Payment    uos.Asset
}
