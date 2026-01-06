package cmd

import (
	"fmt"
	"runtime"
	"react-cli/internal/version"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("react-cli version %s\n", version.Version)
			fmt.Printf("Build time: %s\n", version.BuildTime)
			fmt.Printf("Git commit: %s\n", version.GitCommit)
			fmt.Printf("Go version: %s\n", runtime.Version())
			fmt.Printf("OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
		},
	}
	return versionCmd
}

