package token

import "github.com/tkblack/uos-go"

func init() {
	uos.RegisterAction(AN("wxbio.token"), ActN("transfer"), Transfer{})
	uos.RegisterAction(AN("wxbio.token"), ActN("issue"), Issue{})
	uos.RegisterAction(AN("wxbio.token"), ActN("create"), Create{})
}
