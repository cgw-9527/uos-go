package system

import "github.com/uoscanada/uos-go"

// NewNonce returns a `nonce` action that lives on the
// `uosio.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `uos-bios` boot process by the
// `uosio.system` contract.
func NewVoteProducer(voter uos.AccountName, proxy uos.AccountName, producers ...uos.AccountName) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("voteproducer"),
		Authorization: []uos.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `uosio.system::voteproducer` action
type VoteProducer struct {
	Voter     uos.AccountName   `json:"voter"`
	Proxy     uos.AccountName   `json:"proxy"`
	Producers []uos.AccountName `json:"producers"`
}
