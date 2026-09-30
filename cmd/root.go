/*
Copyright © 2026 Pratyay360

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"

	"github.com/pratyay360/toolbox-export/utils"
	"github.com/spf13/cobra"
)

var logLevel string

// rootCmd runs both exports by default, mirroring the old toolbox-export.sh.
var rootCmd = &cobra.Command{
	Use:   "toolbox-export <application>",
	Short: "Export applications from a toolbox container to the host",
	Long: `toolbox-export exports applications from a toolbox (or any other
container) so they can be launched from the host.

Running it without a subcommand exports both the desktop launcher and the
binary wrapper for the given application.`,
	Args: cobra.ArbitraryArgs,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		utils.SetupLogger(os.Stderr, logLevel)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		app := firstArg(args)
		if err := utils.ExportLauncher(app); err != nil {
			utils.Error(err)
		}
		if err := utils.ExportBinary(app); err != nil {
			utils.Error(err)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.Error(err)
		os.Exit(1)
	}
}

func firstArg(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "log level (trace, debug, info, warn, error)")
	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(binaryCmd)
}
