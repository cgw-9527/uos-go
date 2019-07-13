package msig

import (
	uos "github.com/uoscanada/uos-go"
)

// NewCancel returns a `cancel` action that lives on the
// `uosio.msig` contract.
func NewCancel(proposer uos.AccountName, proposalName uos.Name, canceler uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: uos.AccountName("uosio.msig"),
		Name:    uos.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []uos.PermissionLevel{
			{Actor: canceler, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     uos.AccountName `json:"proposer"`
	ProposalName uos.Name        `json:"proposal_name"`
	Canceler     uos.AccountName `json:"canceler"`
}
