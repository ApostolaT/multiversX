package main

import (
	"fmt"
	"github.com/ApostolaT/multiversX/cmd"
	logger "github.com/multiversx/mx-chain-logger-go"
	"github.com/spf13/cobra"
	"os"
)

var (
	_ = logger.SetLogLevel("*:DEBUG")

	rootCmd = &cobra.Command{
		Use:   "console",
		Short: "Entrypoint for MultiversX console commands",
		Long:  `Commands are ran from the ./bin/multiversX executable`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("MultiversX console command app. Run ./bin/multiversX --help for more info")
		},
	}
)

func main() {
	rootCmd.AddCommand(cmd.IssueNFT)
	rootCmd.AddCommand(cmd.GetTokenID)
	rootCmd.AddCommand(cmd.SetCreatorRole)
	rootCmd.AddCommand(cmd.CreateNFT)
	rootCmd.AddCommand(cmd.SendNFT)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
