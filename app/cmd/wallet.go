/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"goke/modules/wallet"
	"time"

	"github.com/spf13/cobra"
)

// walletCmd represents the wallet command
var walletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		walletRun()
	},
}

func init() {
	rootCmd.AddCommand(walletCmd)
}


func walletRun(){
	dave := wallet.NewAccount("Dave")
	alice := wallet.NewAccount("Alice")

	// Anonymous struct
	movie1 := wallet.NewMovie("freedom", time.Minute * time.Duration(120))

	dave.StoreItem("freedom", movie1)
	dave.TransferTo("freedom", alice)

	fmt.Printf("%+v", alice.GetItem("freedom"))
}