package msig

import (
	uos "github.com/uoscanada/uos-go"
)

// NewApprove returns a `approve` action that lives on the
// `uosio.msig` contract.
func NewApprove(proposer uos.AccountName, proposalName uos.Name, level uos.PermissionLevel) *uos.Action {
	return &uos.Action{
		Account:       uos.AccountName("uosio.msig"),
		Name:          uos.ActionName("approve"),
		Authorization: []uos.PermissionLevel{level},
		ActionData:    uos.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     uos.AccountName     `json:"proposer"`
	ProposalName uos.Name            `json:"proposal_name"`
	Level        uos.PermissionLevel `json:"level"`
}
