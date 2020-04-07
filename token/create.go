package token

import uos "github.com/lialvin/uos-go"

func NewCreate(issuer uos.AccountName, maxSupply uos.Asset) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio.token"),
		Name:    ActN("create"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("wxbio.token"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `wxbio.token` contract.
type Create struct {
	Issuer        uos.AccountName `json:"issuer"`
	MaximumSupply uos.Asset       `json:"maximum_supply"`
}
