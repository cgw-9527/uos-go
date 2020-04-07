package uidsys

import uos "github.com/tkblack/uos-go"

func NewRewardsys(newuid uos.AccountName, recommenduid uint64) *uos.Action {
	return &uos.Action{
		Account: AN("uosuidwallet"),
		Name:    ActN("rewardsys"),
		Authorization: []uos.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(Rewardsys{
			Newuid:     newuid,
			Recommenduid:       recommenduid,
		}),
	}
}

// Transfer represents the `transfer` struct on `wxbio.token` contract.
type Rewardsys struct {
	Newuid     uos.AccountName `json:"newuid"`
	Recommenduid      uint64  `json:"recommenduid"`
}

