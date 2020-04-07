package system

import "github.com/lialvin/uos-go"

// NewCancelDelay creates an action from the `wxbio.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth uos.PermissionLevel, transactionID uos.Checksum256) *uos.Action {
	a := &uos.Action{
		Account: AN("wxbio"),
		Name:    ActN("canceldelay"),
		Authorization: []uos.PermissionLevel{
			cancelingAuth,
		},
		ActionData: uos.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth uos.PermissionLevel `json:"canceling_auth"`
	TransactionID uos.Checksum256     `json:"trx_id"`
}
