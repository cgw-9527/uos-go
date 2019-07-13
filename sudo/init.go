package sudo

import uos "github.com/uoscanada/uos-go"

func init() {
	uos.RegisterAction(AN("uosio.wrap"), ActN("exec"), Exec{})
}

var AN = uos.AN
var ActN = uos.ActN
