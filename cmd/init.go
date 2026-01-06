package cmd

import (
	"fmt"
	"react-cli/internal/config"

	"github.com/spf13/cobra"
)

func InitCmd() *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize react-cli.json configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config.GetDefaultConfig()
			err := config.WriteConfig(cfg)
			if err != nil {
				fmt.Println("Error writing config file:", err)
				return
			}
			fmt.Println("react-cli.json configuration file created successfully")
		},
	}
	return initCmd
}
