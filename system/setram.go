package system

import (
	uos "github.com/lialvin/uos-go"
)

func NewSetRAM(maxRAMSize uint64) *uos.Action {
	a := &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("setram"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      AN("uosio"),
				Permission: uos.PermissionName("active"),
			},
		},
		ActionData: uos.NewActionData(SetRAM{
			MaxRAMSize: uos.Uint64(maxRAMSize),
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize uos.Uint64 `json:"max_ram_size"`
}
