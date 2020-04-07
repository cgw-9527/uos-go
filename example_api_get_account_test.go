package uos_test

import (
	"encoding/json"
	"fmt"

	uos "github.com/tkblack/uos-go"
)

func ExampleAPI_GetAccount() {
	api := uos.New(getAPIURL())

	account := uos.AccountName("uos.rex")
	info, err := api.GetAccount(account)
	if err != nil {
		if err == uos.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
