package msig

import (
	"github.com/uoscanada/uos-go"
)

func init() {
	uos.RegisterAction(AN("uosio.msig"), ActN("propose"), &Propose{})
	uos.RegisterAction(AN("uosio.msig"), ActN("approve"), &Approve{})
	uos.RegisterAction(AN("uosio.msig"), ActN("unapprove"), &Unapprove{})
	uos.RegisterAction(AN("uosio.msig"), ActN("cancel"), &Cancel{})
	uos.RegisterAction(AN("uosio.msig"), ActN("exec"), &Exec{})
}

var AN = uos.AN
var PN = uos.PN
var ActN = uos.ActN
