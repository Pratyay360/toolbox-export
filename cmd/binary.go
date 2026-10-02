/*
Copyright © 2026 Pratyay360 <pratyaymustafi@outlook.com>
*/
package cmd

import (
	"github.com/Pratyay360/toolbox-export/utils"
	"github.com/spf13/cobra"
)

var binaryCmd = &cobra.Command{
	Use:     "binary <application>",
	Aliases: []string{"bexport"},
	Short:   "Export a wrapper script for an application's binary",
	Args:    cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.ExportBinary(firstArg(args))
	},
}
