package rex

import (
	uos "github.com/lialvin/uos-go"
)

func NewDefundCPULoan(from uos.AccountName, loanNumber uint64, amount uos.Asset) *uos.Action {
	return &uos.Action{
		Account: REXAN,
		Name:    ActN("defcpuloan"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(DefundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundCPULoan struct {
	From       uos.AccountName
	LoanNumber uint64
	Amount     uos.Asset
}
