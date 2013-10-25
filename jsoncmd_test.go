// Copyright (c) 2013 Conformal Systems LLC.
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// this has to be in the real json subpackage so we can mock up structs
package btcjson

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

var jsoncmdtests = []struct {
	name   string
	f      func() (Cmd, error)
	result Cmd // after marshal and unmarshal
}{
	{
		name: "basic addmultisigaddress",
		f: func() (Cmd, error) {
			return NewAddMultisigAddressCmd(float64(1), 1,
				[]string{"foo", "bar"})
		},
		result: &AddMultisigAddressCmd{
			id:        float64(1),
			NRequired: 1,
			Keys:      []string{"foo", "bar"},
			Account:   "",
		},
	},
	{
		name: "addmultisigaddress + optional",
		f: func() (Cmd, error) {
			return NewAddMultisigAddressCmd(float64(1), 1,
				[]string{"foo", "bar"}, "address")
		},
		result: &AddMultisigAddressCmd{
			id:        float64(1),
			NRequired: 1,
			Keys:      []string{"foo", "bar"},
			Account:   "address",
		},
	},
	// TODO(oga) Too many arguments to newaddmultisigaddress
	{
		name: "basic addnode add",
		f: func() (Cmd, error) {
			return NewAddNodeCmd(float64(1), "address",
				"add")
		},
		result: &AddNodeCmd{
			id:     float64(1),
			Addr:   "address",
			SubCmd: "add",
		},
	},
	{
		name: "basic addnode remoe",
		f: func() (Cmd, error) {
			return NewAddNodeCmd(float64(1), "address",
				"remove")
		},
		result: &AddNodeCmd{
			id:     float64(1),
			Addr:   "address",
			SubCmd: "remove",
		},
	},
	{
		name: "basic addnode onetry",
		f: func() (Cmd, error) {
			return NewAddNodeCmd(float64(1), "address",
				"onetry")
		},
		result: &AddNodeCmd{
			id:     float64(1),
			Addr:   "address",
			SubCmd: "onetry",
		},
	},
	// TODO(oga) try invalid subcmds
	{
		name: "basic backupwallet",
		f: func() (Cmd, error) {
			return NewBackupWalletCmd(float64(1), "destination")
		},
		result: &BackupWalletCmd{
			id:          float64(1),
			Destination: "destination",
		},
	},
	{
		name: "basic createmultisig",
		f: func() (Cmd, error) {
			return NewCreateMultisigCmd(float64(1), 1,
				[]string{"key1", "key2", "key3"})
		},
		result: &CreateMultisigCmd{
			id:        float64(1),
			NRequired: 1,
			Keys:      []string{"key1", "key2", "key3"},
		},
	},
	{
		name: "basic createrawtransaction",
		f: func() (Cmd, error) {
			return NewCreateRawTransactionCmd(float64(1),
				[]TransactionInput{
					TransactionInput{Txid: "tx1", Vout: 1},
					TransactionInput{Txid: "tx2", Vout: 3}},
				map[string]int64{"bob": 1, "bill": 2})
		},
		result: &CreateRawTransactionCmd{
			id: float64(1),
			Inputs: []TransactionInput{
				TransactionInput{Txid: "tx1", Vout: 1},
				TransactionInput{Txid: "tx2", Vout: 3},
			},
			Amounts: map[string]int64{
				"bob":  1,
				"bill": 2,
			},
		},
	},
	{
		name: "basic decoderawtransaction",
		f: func() (Cmd, error) {
			return NewDecodeRawTransactionCmd(float64(1),
				"thisisahexidecimaltransaction")
		},
		result: &DecodeRawTransactionCmd{
			id: float64(1),
			HexTx: "thisisahexidecimaltransaction",
		},
	},
	{
		name: "basic dumpprivkey",
		f: func() (Cmd, error) {
			return NewDumpPrivKeyCmd(float64(1),
				"address")
		},
		result: &DumpPrivKeyCmd{
			id: float64(1),
			Address: "address",
		},
	},
	{
		name: "basic dumpwallet",
		f: func() (Cmd, error) {
			return NewDumpWalletCmd(float64(1),
				"filename")
		},
		result: &DumpWalletCmd{
			id: float64(1),
			Filename: "filename",
		},
	},
	{
		name: "basic encryptwallet",
		f: func() (Cmd, error) {
			return NewEncryptWalletCmd(float64(1),
				"passphrase")
		},
		result: &EncryptWalletCmd{
			id: float64(1),
			Passphrase: "passphrase",
		},
	},
	{
		name: "basic getaccount",
		f: func() (Cmd, error) {
			return NewGetAccountCmd(float64(1),
				"address")
		},
		result: &GetAccountCmd{
			id: float64(1),
			Address: "address",
		},
	},
	{
		name: "basic getaccountaddress",
		f: func() (Cmd, error) {
			return NewGetAccountAddressCmd(float64(1),
				"account")
		},
		result: &GetAccountAddressCmd{
			id: float64(1),
			Account: "account",
		},
	},
	{
		name: "basic ping",
		f: func() (Cmd, error) {
			return NewPingCmd(float64(1))
		},
		result: &PingCmd{
			id: float64(1),
		},
	},
	{
		name: "basic getblockcount",
		f: func() (Cmd, error) {
			return NewGetBlockCountCmd(float64(1))
		},
		result: &GetBlockCountCmd{
			id: float64(1),
		},
	},
	{
		name: "basic getblock",
		f: func() (Cmd, error) {
			return NewGetBlockCmd(float64(1),
				"somehash")
		},
		result: &GetBlockCmd{
			id:   float64(1),
			Hash: "somehash",
		},
	},
}

func TestCmds(t *testing.T) {
	for _, test := range jsoncmdtests {
		c, err := test.f()
		if err != nil {
			t.Errorf("%s: failed to run func: %v",
				test.name, err)
			continue
		}

		msg, err := json.Marshal(c)
		if err != nil {
			t.Errorf("%s: failed to marshal cmd: %v",
				test.name, err)
			continue
		}

		c2, err := ParseMarshaledCmd(msg)
		if err != nil {
			t.Errorf("%s: failed to ummarshal cmd: %v",
				test.name, err)
			continue
		}

		if !reflect.DeepEqual(test.result, c2) {
			t.Errorf("%s: unmarshal not as expected. "+
				"got %v wanted %v", test.name, spew.Sdump(c2),
				spew.Sdump(test.result))
		}
		if !reflect.DeepEqual(c, c2) {
			t.Errorf("%s: unmarshal not as we started with. "+
				"got %v wanted %v", test.name, spew.Sdump(c2),
				spew.Sdump(c))
		}

	}
}
