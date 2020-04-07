package sudo

import uos "github.com/tkblack/uos-go"

func init() {
	uos.RegisterAction(AN("wxbio.wrap"), ActN("exec"), Exec{})
}

var AN = uos.AN
var ActN = uos.ActN
