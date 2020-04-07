package system

import (
	uos "github.com/lialvin/uos-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `wxbio.system` contract.
func NewRegProxy(proxy uos.AccountName, isProxy bool) *uos.Action {
	return &uos.Action{
		Account: AN("wxbio"),
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

// RegProxy represents the `wxbio.system::regproxy` action
type RegProxy struct {
	Proxy   uos.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}
