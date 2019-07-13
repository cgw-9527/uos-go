package forum

import uos "github.com/uoscanada/uos-go"

func init() {
	uos.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	uos.RegisterAction(ForumAN, ActN("expire"), Expire{})
	uos.RegisterAction(ForumAN, ActN("post"), Post{})
	uos.RegisterAction(ForumAN, ActN("propose"), Propose{})
	uos.RegisterAction(ForumAN, ActN("status"), Status{})
	uos.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	uos.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	uos.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = uos.AN
var PN = uos.PN
var ActN = uos.ActN

var ForumAN = AN("uosio.forum")
