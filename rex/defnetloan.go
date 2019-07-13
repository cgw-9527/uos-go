package rex

import (
	uos "github.com/uoscanada/uos-go"
)

func NewDefundNetLoan(from uos.AccountName, loanNumber uint64, amount uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("defnetloan"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(DefundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundNetLoan struct {
	From       uos.AccountName
	LoanNumber uint64
	Amount     uos.Asset
}
