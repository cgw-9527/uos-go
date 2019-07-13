package uos_test

import (
	"encoding/json"
	"fmt"

	uos "github.com/uoscanada/uos-go"
)

func ExampleAPI_GetInfo() {
	api := uos.New(getAPIURL())

	info, err := api.GetInfo()
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
