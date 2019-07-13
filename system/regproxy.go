package system

import (
	uos "github.com/uoscanada/uos-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `uosio.system` contract.
func NewRegProxy(proxy uos.AccountName, isProxy bool) *uos.Action {
	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("regproxy"),
		Authorization: []uos.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: uos.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `uosio.system::regproxy` action
type RegProxy struct {
	Proxy   uos.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}
