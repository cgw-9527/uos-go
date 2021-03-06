package msig

import (
	uos "github.com/tkblack/uos-go"
)

// NewExec returns a `exec` action that lives on the
// `wxbio.msig` contract.
func NewExec(proposer uos.AccountName, proposalName uos.Name, executer uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: uos.AccountName("wxbio.msig"),
		Name:    uos.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []uos.PermissionLevel{
			{Actor: executer, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     uos.AccountName `json:"proposer"`
	ProposalName uos.Name        `json:"proposal_name"`
	Executer     uos.AccountName `json:"executer"`
}
