package uidsys

import uos "github.com/lialvin/uos-go"

func NewSetdata(key uos.AccountName, val  uint64 , op uint8 ) *uos.Action {
	return &uos.Action{
		Account: AN("uosuidwallet"),
		Name:    ActN("setdata"),
		Authorization: []uos.PermissionLevel{
			{Actor: AN("uosuidwallet"), Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Setdata{
			Key:        key,
			Val:        val,
			Op:         op,
		}),
	}
}

// Create represents the `create` struct on the `uosio.token` contract.
type Setdata struct {
	Key     uos.AccountName `json:"key"`
	Val     uint64       `json:"val"`
	Op      uint8        `json:"val"`
}


