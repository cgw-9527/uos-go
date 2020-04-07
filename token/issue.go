package token

import uos "github.com/lialvin/uos-go"

func NewIssue(to uos.AccountName, quantity uos.Asset, memo string) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio.token"),
		Name:    ActN("issue"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("wxbio"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `wxbio.token` contract.
type Issue struct {
	To       uos.AccountName `json:"to"`
	Quantity uos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
