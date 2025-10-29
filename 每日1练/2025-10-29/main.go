package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version    = "v1.0"
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Prints the version of tacoctl",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("tacoctl version:", version)
		},
	}
)

var rootCmd = &cobra.Command{
	Use:   "leetcode-go",
	Short: "A simple command line client for leetcode-go.",
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	rootCmd.AddCommand(
		versionCmd,
	)
}

func main() {
	execute()
}
