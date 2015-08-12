package main

import (
	"fmt"
	"io/ioutil"

	"github.com/eris-ltd/mint-client/mintx/core"

	"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/spf13/cobra"
)

// do we really need these?
/*
func cliInput(cmd *cobra.Command, args []string) {
	input, err := coreInput(pubkey, amtS, nonceS, addr)
	common.IfExit(err)
	fmt.Printf("%s\n", input)
}

func cliOutput(cmd *cobra.Command, args []string) {
	output, err := coreOutput(addr, amtS)
	common.IfExit(err)
	fmt.Printf("%s\n", output)
}
*/

func cliSend(cmd *cobra.Command, args []string) {
	tx, err := core.Send(nodeAddr, pubkey, addr, toAddr, amtS, nonceS)
	common.IfExit(err)
	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func cliName(cmd *cobra.Command, args []string) {
	if data != "" && dataFile != "" {
		common.Exit(fmt.Errorf("Please specify only one of --data and --data-file"))
	}
	if data == "" && dataFile != "" {
		b, err := ioutil.ReadFile(dataFile)
		common.IfExit(err)
		data = string(b)
	}
	tx, err := core.Name(nodeAddr, pubkey, addr, amtS, nonceS, feeS, name, data)
	common.IfExit(err)
	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func cliCall(cmd *cobra.Command, args []string) {
	tx, err := core.Call(nodeAddr, pubkey, addr, toAddr, amtS, nonceS, gasS, feeS, data)
	common.IfExit(err)
	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func cliPermissions(cmd *cobra.Command, args []string) {
	// all functions take at least 2 args (+ name)
	if len(args) < 3 {
		common.Exit(fmt.Errorf("Please enter the permission function you'd like to call, followed by it's arguments"))
	}
	permFunc := args[0]
	tx, err := core.Permissions(nodeAddr, pubkey, addr, nonceS, permFunc, args[1:])
	common.IfExit(err)
	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func cliNewAccount(cmd *cobra.Command, args []string) {
	/*

		tx, err := coreNewAccount(nodeAddr,signAddr, pubkey, chainID)
		common.IfExit(err)

		logger.Debugf("%v\n", tx)
		unpackSignAndBroadcast(core.SignAndBroadcast( chainID, nodeAddr,signAddr, tx, sign, broadcast, wait)
	*/
}

func cliBond(cmd *cobra.Command, args []string) {
	tx, err := core.Bond(nodeAddr, pubkey, unbondAddr, amtS, nonceS)
	common.IfExit(err)

	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func cliUnbond(cmd *cobra.Command, args []string) {
	tx, err := core.Unbond(addr, height)
	common.IfExit(err)
	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func cliRebond(cmd *cobra.Command, args []string) {
	tx, err := core.Rebond(addr, height)
	common.IfExit(err)
	logger.Debugf("%v\n", tx)
	unpackSignAndBroadcast(core.SignAndBroadcast(chainID, nodeAddr, signAddr, tx, sign, broadcast, wait))
}

func unpackSignAndBroadcast(result *core.TxResult, err error) {
	common.IfExit(err)
	if result == nil {
		// if we don't provide --sign or --broadcast
		return
	}
	fmt.Printf("Transaction Hash: %X\n", result.Hash)
	if result.Return != nil {
		fmt.Printf("Return Value: %X\n", result.Return)
		fmt.Printf("Exception: %s\n", result.Exception)
	}
}
