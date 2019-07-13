package forum

import (
	uos "github.com/lialvin/uos-go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter uos.AccountName, proposalName uos.Name, voteValue uint8, voteJSON string) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []uos.PermissionLevel{
			{Actor: voter, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `uosio.forum::vote` action.
type Vote struct {
	Voter        uos.AccountName `json:"voter"`
	ProposalName uos.Name        `json:"proposal_name"`
	Vote         uint8           `json:"vote"`
	VoteJSON     string          `json:"vote_json"`
}
