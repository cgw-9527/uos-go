package sudo

import (
	uos "github.com/uoscanada/uos-go"
)

// NewExec creates an `exec` action, found in the `uosio.wrap`
// contract.
//
// Given an `uos.Transaction`, call `uos.MarshalBinary` on it first,
// pass the resulting bytes as `uos.HexBytes` here.
func NewExec(executer uos.AccountName, transaction uos.Transaction) *uos.Action {
	a := &uos.Action{
		Account: uos.AccountName("uosio.wrap"),
		Name:    uos.ActionName("exec"),
		Authorization: []uos.PermissionLevel{
			{Actor: executer, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `uosio.system::exec` action.
type Exec struct {
	Executer    uos.AccountName `json:"executer"`
	Transaction uos.Transaction `json:"trx"`
}
