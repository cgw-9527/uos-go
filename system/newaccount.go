package system

import (
	"github.com/tkblack/uos-go"
	"github.com/tkblack/uos-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `wxbio.system` contract.
func NewNewAccount(creator, newAccount uos.AccountName, publicKey ecc.PublicKey) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("newaccount"),
		Authorization: []uos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: uos.Authority{
				Threshold: 1,
				Keys: []uos.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []uos.PermissionLevelWeight{},
			},
			Active: uos.Authority{
				Threshold: 1,
				Keys: []uos.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []uos.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `wxbio.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount uos.AccountName, delegatedTo uos.AccountName) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("newaccount"),
		Authorization: []uos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: uos.Authority{
				Threshold: 1,
				Keys:      []uos.KeyWeight{},
				Accounts: []uos.PermissionLevelWeight{
					uos.PermissionLevelWeight{
						Permission: uos.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: uos.Authority{
				Threshold: 1,
				Keys:      []uos.KeyWeight{},
				Accounts: []uos.PermissionLevelWeight{
					uos.PermissionLevelWeight{
						Permission: uos.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `wxbio.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount uos.AccountName, owner, active uos.Authority) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("newaccount"),
		Authorization: []uos.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `wxbio.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator uos.AccountName `json:"creator"`
	Name    uos.AccountName `json:"name"`
	Owner   uos.Authority   `json:"owner"`
	Active  uos.Authority   `json:"active"`
}
