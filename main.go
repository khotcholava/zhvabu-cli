package main

import (
	"fmt"
	"react-cli/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "rc",
		Short: "React CLI - Generate React component",
	}

	// Get the generate command from cmd package
	generateCmd := cmd.GenerateCmd()
	rootCmd.AddCommand(generateCmd)

	initCmd := cmd.InitCmd()
	rootCmd.AddCommand(initCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}
