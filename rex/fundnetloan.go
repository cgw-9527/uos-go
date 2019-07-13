package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewFundNetLoan(from uos.AccountName, loanNumber uint64, payment uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("fundnetloan"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(FundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundNetLoan struct {
	From       uos.AccountName
	LoanNumber uint64
	Payment    uos.Asset
}
