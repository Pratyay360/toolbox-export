/*
Copyright © 2026 Pratyay360<pratyaymustafi@outlook.com>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const toolboxPathConfig = `
# toolbox-export: add exported binaries to PATH
case ":$PATH:" in
    *":$HOME/.local/toolbox:"*) ;;
    *) export PATH="$PATH:$HOME/.local/toolbox" ;;
esac
`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the toolbox binary directory and shell PATH",
	Long:  `Create ~/.local/toolbox and add it to PATH in ~/.bashrc, ~/.zshrc, and ~/.profile.`,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("resolve home directory: %w", err)
		}

		folder := filepath.Join(home, ".local", "toolbox")
		info, err := os.Stat(folder)
		if err == nil {
			if !info.IsDir() {
				return fmt.Errorf("%s already exists and is not a directory", folder)
			}
			cmd.Printf("Directory already exists: %s\n", folder)
		} else {
			if !os.IsNotExist(err) {
				return fmt.Errorf("check toolbox directory: %w", err)
			}
			if err := os.MkdirAll(folder, 0o755); err != nil {
				return fmt.Errorf("create toolbox directory: %w", err)
			}
			cmd.Printf("Created directory: %s\n", folder)
		}

		for _, name := range []string{".bashrc", ".zshrc", ".profile"} {
			rcPath := filepath.Join(home, name)
			updated, err := addToolboxPath(rcPath)
			if err != nil {
				return err
			}
			if updated {
				cmd.Printf("Added toolbox PATH configuration to: %s\n", rcPath)
			} else {
				cmd.Printf("Toolbox PATH configuration already exists in: %s\n", rcPath)
			}
		}
		cmd.Println("Open a new shell to apply the PATH changes.")
		return nil
	},
}

func addToolboxPath(rcPath string) (bool, error) {
	data, err := os.ReadFile(rcPath)
	if err != nil && !os.IsNotExist(err) {
		return false, fmt.Errorf("read shell startup file %s: %w", rcPath, err)
	}
	if strings.Contains(string(data), toolboxPathConfig) {
		return false, nil
	}

	addition := toolboxPathConfig
	if len(data) > 0 && data[len(data)-1] != '\n' {
		addition = "\n" + addition
	}
	file, err := os.OpenFile(rcPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return false, fmt.Errorf("open shell startup file %s: %w", rcPath, err)
	}
	_, writeErr := file.WriteString(addition)
	closeErr := file.Close()
	if writeErr != nil {
		return false, fmt.Errorf("write shell startup file %s: %w", rcPath, writeErr)
	}
	if closeErr != nil {
		return false, fmt.Errorf("close shell startup file %s: %w", rcPath, closeErr)
	}
	return true, nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
