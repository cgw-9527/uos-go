package msig

import (
	"github.com/tkblack/uos-go"
)

func init() {
	uos.RegisterAction(AN("wxbio.msig"), ActN("propose"), &Propose{})
	uos.RegisterAction(AN("wxbio.msig"), ActN("approve"), &Approve{})
	uos.RegisterAction(AN("wxbio.msig"), ActN("unapprove"), &Unapprove{})
	uos.RegisterAction(AN("wxbio.msig"), ActN("cancel"), &Cancel{})
	uos.RegisterAction(AN("wxbio.msig"), ActN("exec"), &Exec{})
}

var AN = uos.AN
var PN = uos.PN
var ActN = uos.ActN
