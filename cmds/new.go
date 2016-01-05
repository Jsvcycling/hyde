/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

var newCommand = &cobra.Command{
	Use:   "new [name]",
	Short: "Generate a new Hyde site.",
	Run:   runNewCommand,
}

func runNewCommand(cmd *cobra.Command, args []string) {
	// TODO
}
