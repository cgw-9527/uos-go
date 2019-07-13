package forum

import (
	uos "github.com/uoscanada/uos-go"
)

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster uos.AccountName, postUUID, content string, replyToPoster uos.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *uos.Action {
	a := &uos.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []uos.PermissionLevel{
			{Actor: poster, Permission: uos.PermissionName("active")},
		},
		ActionData: uos.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `uosio.forum::post` action.
type Post struct {
	Poster          uos.AccountName `json:"poster"`
	PostUUID        string          `json:"post_uuid"`
	Content         string          `json:"content"`
	ReplyToPoster   uos.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string          `json:"reply_to_post_uuid"`
	Certify         bool            `json:"certify"`
	JSONMetadata    string          `json:"json_metadata"`
}
