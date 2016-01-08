/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

// TODO: Add flag handling
var generateCommand = &cobra.Command{
	Use:   "generate [type] [name]",
	Short: "Generate a new typed resource integrated it properly.",
	Run:   runGenerateCommand,
}

func runGenerateCommand(cmd *cobra.Command, args []string) {
	// TODO
}
