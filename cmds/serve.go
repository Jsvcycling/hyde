/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

// TODO: Add flag handling
var serveCommand = &cobra.Command{
	Use:   "serve [port]",
	Short: "Use Hyde to serve your website (not recommended for production).",
	Run:   runServeCommand,
}

func runServeCommand(cmd *cobra.Command, args []string) {
	// TODO
}
