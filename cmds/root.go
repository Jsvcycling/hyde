/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

// TODO: Add flag handling
var rootCommand = &cobra.Command{
	Use:   "hyde",
	Short: "Hyde is a lightweight, static page generator written in Go.",
}

func Run() {
	rootCommand.AddCommand(buildCommand)
	rootCommand.AddCommand(newCommand)
	rootCommand.AddCommand(serveCommand)
}
