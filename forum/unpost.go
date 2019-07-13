package forum

import (
	uos "github.com/uoscanada/uos-go"
)

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster uos.AccountName, postUUID string) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("unpost"),
		Authorization: []uos.PermissionLevel{
			{Actor: poster, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `uosio.forum::unpost` action.
type UnPost struct {
	Poster   uos.AccountName `json:"poster"`
	PostUUID string          `json:"post_uuid"`
}
