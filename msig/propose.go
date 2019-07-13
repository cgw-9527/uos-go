package msig

import (
	uos "github.com/lialvin/uos-go"
)

// NewPropose returns a `propose` action that lives on the
// `uosio.msig` contract.
func NewPropose(proposer uos.AccountName, proposalName uos.Name, requested []uos.PermissionLevel, transaction *uos.Transaction) *uos.Action {
	return &uos.Action{
		Account: uos.AccountName("uosio.msig"),
		Name:    uos.ActionName("propose"),
		Authorization: []uos.PermissionLevel{
			{Actor: proposer, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     uos.AccountName       `json:"proposer"`
	ProposalName uos.Name              `json:"proposal_name"`
	Requested    []uos.PermissionLevel `json:"requested"`
	Transaction  *uos.Transaction      `json:"trx"`
}
