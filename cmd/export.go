/*
Copyright © 2026 Pratyay360 <pratyaymustafi@outlook.com>
*/
package cmd

import (
	"github.com/pratyay360/toolbox-export/utils"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export <application>",
	Short: "Export an application's desktop launcher and icons",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.ExportLauncher(firstArg(args))
	},
}
