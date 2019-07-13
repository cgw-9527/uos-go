package system

import (
	uos "github.com/uoscanada/uos-go"
)

func NewSetRAMRate(bytesPerBlock uint16) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("setram"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      AN("uosio"),
				Permission: uos.PermissionName("active"),
			},
		},
		ActionData: uos.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}
