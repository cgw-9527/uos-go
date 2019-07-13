package token

import uos "github.com/uoscanada/uos-go"

func NewIssue(to uos.AccountName, quantity uos.Asset, memo string) *uos.Action {
	return &uos.Action{
		Account: AN("uosio.token"),
		Name:    ActN("issue"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("uosio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `uosio.token` contract.
type Issue struct {
	To       uos.AccountName `json:"to"`
	Quantity uos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
