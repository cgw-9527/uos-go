package forum

import (
	uos "github.com/uoscanada/uos-go"
)

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter uos.AccountName, proposalName uos.Name) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []uos.PermissionLevel{
			{Actor: voter, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `uosio.forum::unvote` action.
type UnVote struct {
	Voter        uos.AccountName `json:"voter"`
	ProposalName uos.Name        `json:"proposal_name"`
}
