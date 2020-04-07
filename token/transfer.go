package token

import uos "github.com/tkblack/uos-go"

func NewTransfer(from, to uos.AccountName, quantity uos.Asset, memo string) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio.token"),
		Name:    ActN("transfer"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `wxbio.token` contract.
type Transfer struct {
	From     uos.AccountName `json:"from"`
	To       uos.AccountName `json:"to"`
	Quantity uos.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}
