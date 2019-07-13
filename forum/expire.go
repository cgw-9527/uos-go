package forum

import (
	uos "github.com/lialvin/uos-go"
)

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer uos.AccountName, proposalName uos.Name) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []uos.PermissionLevel{
			{Actor: proposer, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `uosio.forum::propose` action.
type Expire struct {
	ProposalName uos.Name `json:"proposal_name"`
}
