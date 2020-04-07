package uidsys

import "github.com/tkblack/uos-go"

func init() {
	uos.RegisterAction(AN("uosuidwallet"), ActN("setdata"), Setdata{})
	uos.RegisterAction(AN("uosuidwallet"), ActN("forcedel"), Forcedel{})	
	uos.RegisterAction(AN("uosuidwallet"), ActN("rewardsys"), Rewardsys{})
}

/*
ACTION  setdata(const name& key, const uint64_t& val, const uint8_t& op);

ACTION  forcedel( const name& account_name, const name& newname );

ACTION  dropuid(const name& account_name, const name& newname ) ;

ACTION  rewardsys(const name& newuid, const uint64_t& recommenduid ) ;
*/