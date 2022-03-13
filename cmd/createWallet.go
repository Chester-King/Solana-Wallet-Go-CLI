package cmd

import (
	"fmt"

	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/spf13/cobra"
)

// createWalletCmd represents the createWallet command
var createWalletCmd = &cobra.Command{
	Use:   "createWallet", 
	Short: "Creates a new wallet", 
	Long: "Creates a new wallet and provides wallet address and private key.",
	
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating new wallet.")
        wallet := CreateNewWallet(rpc.DevnetRPCEndpoint)
        fmt.Println("Public Key: " + wallet.account.PublicKey.ToBase58())
        fmt.Println("Private Key Saved in 'key_data' file")
	},
}

func init() {
	rootCmd.AddCommand(createWalletCmd)

}
