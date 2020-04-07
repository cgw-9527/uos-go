package uidsys

import uos "github.com/tkblack/uos-go"

func NewForcedel(account_name uos.AccountName, newname uos.AccountName ) *uos.Action {
	return &uos.Action{
		Account: AN("uosuidwallet"),
		Name:    ActN("forcedel"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("uosuidwallet"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Forcedel{
			Account_name:       account_name,
			Newname:      newname,			
		}),
	}
}

// Issue represents the `issue` struct on the `wxbio.token` contract.
type Forcedel struct {
	Account_name     uos.AccountName    `json:"account_name"`
	Newname          uos.AccountName      `json:"newname"`	
}

