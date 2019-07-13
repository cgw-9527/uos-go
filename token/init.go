package token

import "github.com/lialvin/uos-go"

func init() {
	uos.RegisterAction(AN("uosio.token"), ActN("transfer"), Transfer{})
	uos.RegisterAction(AN("uosio.token"), ActN("issue"), Issue{})
	uos.RegisterAction(AN("uosio.token"), ActN("create"), Create{})
}
