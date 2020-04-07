package forum

import (
	uos "github.com/lialvin/uos-go"
)

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer uos.AccountName, proposalName uos.Name, title string, proposalJSON string, expiresAt uos.JSONTime) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []uos.PermissionLevel{
			{Actor: proposer, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `wxbio.forum::propose` action.
type Propose struct {
	Proposer     uos.AccountName `json:"proposer"`
	ProposalName uos.Name        `json:"proposal_name"`
	Title        string          `json:"title"`
	ProposalJSON string          `json:"proposal_json"`
	ExpiresAt    uos.JSONTime    `json:"expires_at"`
}
