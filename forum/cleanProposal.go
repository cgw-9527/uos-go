package forum

import (
	uos "github.com/uoscanada/uos-go"
)

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner uos.AccountName, proposalName uos.Name, maxCount uint64) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []uos.PermissionLevel{
			{Actor: cleaner, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `uosio.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName uos.Name `json:"proposal_name"`
	MaxCount     uint64   `json:"max_count"`
}
