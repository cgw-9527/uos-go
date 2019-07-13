package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	uos "github.com/lialvin/uos-go"
)

func NewSetContract(account uos.AccountName, wasmPath, abiPath string) (out []*uos.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*uos.Action{codeAction, abiAction}, nil
}

func NewSetCode(account uos.AccountName, wasmPath string) (out *uos.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("setcode"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      account,
				Permission: uos.PermissionName("active"),
			},
		},
		ActionData: uos.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      uos.HexBytes(codeContent),
		}),
	}, nil
}

func NewSetABI(account uos.AccountName, abiPath string) (out *uos.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef uos.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %s", err)
		}

		abiPacked, err = uos.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %s", err)
		}
	}

	return &uos.Action{
		Account: AN("uosio"),
		Name:    ActN("setabi"),
		Authorization: []uos.PermissionLevel{
			{
				Actor:      account,
				Permission: uos.PermissionName("active"),
			},
		},
		ActionData: uos.NewActionData(SetABI{
			Account: account,
			ABI:     uos.HexBytes(abiPacked),
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account uos.AccountName, wasmPath, abiPath string) (out *uos.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &uos.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   uos.AccountName `json:"account"`
	VMType    byte            `json:"vmtype"`
	VMVersion byte            `json:"vmversion"`
	Code      uos.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account uos.AccountName `json:"account"`
	ABI     uos.HexBytes    `json:"abi"`
}
