package cmd

import (
	"fmt"
	"os"
	"react-cli/internal/config"
	"react-cli/internal/generator"
	"react-cli/internal/parser"
	"strings"

	"github.com/spf13/cobra"
)

func GenerateCmd() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate React component",
	}

	generateCmd.AddCommand(ComponentCmd())
	return generateCmd
}

func ComponentCmd() *cobra.Command {
	componentCmd := &cobra.Command{
		Use:   "component",
		Short: "Generate React component",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Error: Component name is required")
				fmt.Println("Usage: rc generate component <name> [\"<props>\"]")
				return
			}

			style, _ := cmd.Flags().GetString("style")
			path, _ := cmd.Flags().GetString("path")
			skipStyle, _ := cmd.Flags().GetBool("skip-style")
			prefix, _ := cmd.Flags().GetString("prefix")
			memo, _ := cmd.Flags().GetBool("memo")
			forwardRef, _ := cmd.Flags().GetBool("forward-ref")
			class, _ := cmd.Flags().GetBool("class")

			// Read config
			cfg, err := config.ReadConfig()
			if err != nil {
				fmt.Println("Error reading config:", err)
				return
			}

			// Merge config with flags (flags override config)
			finalStyle := cfg.Defaults.Component.Style
			if style != "" {
				finalStyle = style
			}

			finalPath := cfg.Defaults.Component.Path
			if path != "" {
				finalPath = path
			}

			finalSkipStyle := cfg.Defaults.Component.SkipStyle
			if skipStyle {
				finalSkipStyle = true
			}

			finalPrefix := cfg.Project.Prefix
			if prefix != "" {
				finalPrefix = prefix
			}

			finalMemo := cfg.Defaults.Component.Memo
			if memo {
				finalMemo = true
			}

			finalForwardRef := cfg.Defaults.Component.ForwardRef
			if forwardRef {
				finalForwardRef = true
			}

			finalClass := cfg.Defaults.Component.Class
			if class {
				finalClass = true
			}

			componentName := args[0]
			finalComponentName := componentName

			// Add prefix to component name if prefix is set
			if finalPrefix != "" {
				finalComponentName = fmt.Sprintf("%s%s",
					strings.ToUpper(string(finalPrefix[0]))+finalPrefix[1:],
					componentName)
			}
			props := ""

			if len(args) >= 2 {
				props = args[1]
			}

			parseProps, err := parser.ParseProps(props)

			if err != nil {
				fmt.Printf("Error parsing props: %v\n", err)
				fmt.Println("Example: rc generate component UserList \"userList, onClick, isActive\"")
				return
			}

			// Build full path
			fullPath := finalPath
			if finalPath != "." && finalPath != "" {
				fullPath = fmt.Sprintf("%s/%s", finalPath, finalComponentName)
			} else {
				fullPath = finalComponentName
			}

			// Generate component content
			componentContent := generator.GenerateComponentContent(finalComponentName, parseProps, finalStyle, finalMemo, finalForwardRef, finalClass)

			// Create directory
			err = os.MkdirAll(fullPath, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}

			// Write component file
			filename := fmt.Sprintf("%s/%s.tsx", fullPath, finalComponentName)
			err = os.WriteFile(filename, []byte(componentContent), 0644)
			if err != nil {
				fmt.Println("Error writing file:", err)
				return
			}
			fmt.Printf("Created %s\n", filename)

			// Create style file (if not skipped)
			if !finalSkipStyle && finalStyle != "none" {
				err = generator.CreateStyleFile(finalComponentName, finalStyle, fullPath)
				if err != nil {
					fmt.Println("Error creating style file:", err)
					return
				}
				fmt.Printf("Created %s.%s\n", finalComponentName, finalStyle)
			}

			// Create index file
			err = generator.CreateIndexFile(finalComponentName, fullPath)
			if err != nil {
				fmt.Println("Error creating index file:", err)
				return
			}
			fmt.Printf("Created index.ts\n")
		},
	}
	componentCmd.Flags().String("style", "", "Style file type (css, scss, sass, none)")
	componentCmd.Flags().String("path", "", "Path where to create component")
	componentCmd.Flags().String("prefix", "", "Prefix to add to component name")
	componentCmd.Flags().Bool("skip-tests", false, "Skip test file generation")
	componentCmd.Flags().Bool("skip-style", false, "Skip style file generation")
	componentCmd.Flags().Bool("memo", false, "Wrap component with React.memo()")
	componentCmd.Flags().Bool("forward-ref", false, "Wrap component with forwardRef()")
	componentCmd.Flags().Bool("class", false, "Generate class component instead of functional component")

	return componentCmd
}
