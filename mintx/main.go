package main

import (
	"fmt"
	"os"

	"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
	"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/spf13/cobra"
)

var (
	DefaultKeyDaemonHost = "localhost"
	DefaultKeyDaemonPort = "4767"
	DefaultKeyDaemonAddr = "http://" + DefaultKeyDaemonHost + ":" + DefaultKeyDaemonPort

	DefaultNodeRPCHost = "pinkpenguin.chaintest.net"
	DefaultNodeRPCPort = "46657"
	DefaultNodeRPCAddr = "http://" + DefaultNodeRPCHost + ":" + DefaultNodeRPCPort + "/"

	DefaultPubKey  string
	DefaultChainID string
)

// override the hardcoded defaults with env variables if they're set

//TODO
func init() {
	signAddr := os.Getenv("MINTX_SIGN_ADDR")
	if signAddr != "" {
		DefaultKeyDaemonAddr = signAddr
	}

	nodeAddr := os.Getenv("MINTX_NODE_ADDR")
	if nodeAddr != "" {
		DefaultNodeRPCAddr = nodeAddr
	}

	pubkey := os.Getenv("MINTX_PUBKEY")
	if pubkey != "" {
		DefaultPubKey = pubkey
	}

	chainID := os.Getenv("MINTX_CHAINID")
	if chainID != "" {
		DefaultChainID = chainID
	}
}

var (

	// flags with env var defaults
	signAddr string
	nodeAddr string
	pubkey   string
	chainID  string

	sign      bool
	broadcast bool
	wait      bool

	// tx flags
	amtS       string
	nonceS     string
	addr       string
	name       string
	data       string
	dataFile   string
	toAddr     string
	feeS       string
	gasS       string
	unbondAddr string
	height     string

	debugFlag bool
)

func main() {

	//------------------------------------------------------------
	// main tx commands

	/* TODO decide if needed
	var inputCmd = &cobra.Command{
				Use:   "input",
				Short:  "mintx input --pubkey <pubkey> --amt <amt> --nonce <nonce>",
				Run: cliInput,
			}

	var outputCmd = &cobra.Command{
				Use:   "output",
				Short:  "mintx output --addr <addr> --amt <amt>",
				Run: cliOutput,
			}
	*/

	var sendCmd = &cobra.Command{
		Use:   "send",
		Short: "mintx send --amt <amt> --to <addr>",
		Run:   cliSend,
	}

	var nameCmd = &cobra.Command{
		Use:   "name",
		Short: "mintx name --amt <amt> --name <name> --data <data>",
		Run:   cliName,
	}

	var callCmd = &cobra.Command{
		Use:   "call",
		Short: "mintx call --amt <amt> --fee <fee> --gas <gas> --to <contract addr> --data <data>",
		Run:   cliCall,
	}

	var bondCmd = &cobra.Command{
		Use:   "bond",
		Short: "mintx bond --pubkey <pubkey> --amt <amt> --unbond-to <address>",
		Run:   cliBond,
	}

	var unbondCmd = &cobra.Command{
		Use:   "unbond",
		Short: "mintx unbond --addr <address> --height <block_height>",
		Run:   cliUnbond,
	}

	var rebondCmd = &cobra.Command{
		Use:   "rebond",
		Short: "mintx rebond --addr <address> --height <block_height>",
		Run:   cliRebond,
	}

	var permissionsCmd = &cobra.Command{
		Use:   "perm",
		Short: "mintx perm <function name> <args ...>",
		Run:   cliPermissions,
	}

	var newAccountCmd = &cobra.Command{
		Use:   "new",
		Short: "mintx new",
		Run:   cliNewAccount,
	}

	//XXX root
	var rootCmd = &cobra.Command{
		Use:   "mintx",
		Short: "Create and broadcast tendermint txs",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var level int
			if debugFlag {
				level = 2
			}
			log.SetLoggers(level, os.Stdout, os.Stderr)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			log.Flush()
		},
	}

	//TODO disentangle non-global flags
	// these are defined in here so we can update the
	// defaults with env variables first
	rootCmd.PersistentFlags().StringVarP(&signAddr, "sign-addr", "", DefaultKeyDaemonAddr, "set the address of the eris-keys daemon")
	rootCmd.PersistentFlags().StringVarP(&nodeAddr, "node-addr", "", DefaultNodeRPCAddr, "set the address of the tendermint rpc server")
	rootCmd.PersistentFlags().StringVarP(&pubkey, "pubkey", "", DefaultPubKey, "specity the pubkey")
	rootCmd.PersistentFlags().StringVarP(&chainID, "chainID", "", DefaultChainID, "specify the chainID")

	//----------------------------------------------------------------
	// optional action flags
	rootCmd.PersistentFlags().BoolVarP(&sign, "sign", "", false, "sign the transaction using the daemon at MINTX_SIGN_ADDR")
	rootCmd.PersistentFlags().BoolVarP(&broadcast, "broadcast", "", false, "broadcast the transaction using the daemon at MINTX_NODE_ADDR")
	rootCmd.PersistentFlags().BoolVarP(&wait, "wait", "", false, "wait for the transaction to be committed in a block")

	//----------------------------------------------------------------
	// tx data flags
	rootCmd.PersistentFlags().StringVarP(&amtS, "amt", "", "", "specify an amount ")
	rootCmd.PersistentFlags().StringVarP(&nonceS, "nonce", "", "", "set the account nonce")
	rootCmd.PersistentFlags().StringVarP(&addr, "addr", "", "", "specify an address")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "", "", "specify a name")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "", "", "specify some data")
	rootCmd.PersistentFlags().StringVarP(&dataFile, "data-file", "", "", "specify a file with some data")
	rootCmd.PersistentFlags().StringVarP(&toAddr, "to", "", "", "specify and address to send to")
	rootCmd.PersistentFlags().StringVarP(&feeS, "fee", "", "", "specify the fee to send")
	rootCmd.PersistentFlags().StringVarP(&gasS, "gas", "", "", "specify the gas limit for a CallTx")
	rootCmd.PersistentFlags().StringVarP(&unbondAddr, "unbond-to", "", "", "specify an address to unbond to")
	rootCmd.PersistentFlags().StringVarP(&height, "height", "", "", "specify a height to unbond at")

	//Formatting Flags
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "print debug messages")

	rootCmd.AddCommand(
		// inputCmd,
		// outputCmd,
		sendCmd,
		nameCmd,
		callCmd,
		bondCmd,
		unbondCmd,
		rebondCmd,
		// dupeoutCmd,
		permissionsCmd,
		newAccountCmd,
	)
	rootCmd.Execute()

}

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
