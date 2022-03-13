package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer", 
	Short: "Transfer SOL", 
	Long: "Transfer SOL from your wallet to other Solana wallets.",
	
	Run: func(cmd *cobra.Command, args []string) {
		 fmt.Println("Recepient address: " + args[0]) 
            fmt.Println("Amount to be sent: " + args[1]) 
            amount, _ := strconv.ParseUint(args[1], 10, 64) 
            txhash, _ := Transfer(args[0], amount) 
            fmt.Println("Transaction complete.\nTransaction hash: " + txhash)
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)

}
