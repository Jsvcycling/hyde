/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

var port int

var serveCommand = &cobra.Command{
	Use:     "serve [port]",
	Aliases: []string{"server"},
	Short:   "Use Hyde to serve your website (not recommended for production).",
	Run:     runServeCommand,
}

func init() {
	serveCommand.Flags().IntVarP(&port, "port", "p", 3000, "The port to use")
}

func runServeCommand(cmd *cobra.Command, args []string) {
	// TODO
}
