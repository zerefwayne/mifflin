package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mifflin",
	Short: "Mifflin is a friendly linux system manager",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Mifflin has started")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
