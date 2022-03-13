package cmd

import (
	"fmt"
	"strconv"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// importWalletCmd represents the importWallet command
var importWalletCmd = &cobra.Command{
	Use:   "importWallet", 
	Short: "Imports and existing wallet", 
	Long: "Imports and existing wallet from a given private key.",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Importing wallet from the 'key_data' file.")
		wallet, _ := ImportOldWallet(rpc.DevnetRPCEndpoint)
		fmt.Println("Public Key: " + wallet.account.PublicKey.ToBase58())
		balance, _ := GetBalance()
		fmt.Println("Wallet balance: " + strconv.Itoa(int(balance/1e9)) + "SOL")
	},
}

func init() {
	rootCmd.AddCommand(importWalletCmd)

	
}
