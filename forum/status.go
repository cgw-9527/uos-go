package forum

import (
	uos "github.com/lialvin/uos-go"
)

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account uos.AccountName, content string) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []uos.PermissionLevel{
			{Actor: account, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `wxbio.forum::status` action.
type Status struct {
	Account uos.AccountName `json:"account_name"`
	Content string          `json:"content"`
}
