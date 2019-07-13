package msig

import (
	uos "github.com/lialvin/uos-go"
)

// NewUnapprove returns a `unapprove` action that lives on the
// `uosio.msig` contract.
func NewUnapprove(proposer uos.AccountName, proposalName uos.Name, level uos.PermissionLevel) *uos.Action {
	return &uos.Action{
		Account:       uos.AccountName("uosio.msig"),
		Name:          uos.ActionName("unapprove"),
		Authorization: []uos.PermissionLevel{level},
		ActionData:    uos.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     uos.AccountName     `json:"proposer"`
	ProposalName uos.Name            `json:"proposal_name"`
	Level        uos.PermissionLevel `json:"level"`
}
