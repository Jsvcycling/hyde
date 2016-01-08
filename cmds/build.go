/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

// TODO: Add flag handling
var buildCommand = &cobra.Command{
	Use:   "build",
	Short: "Build a set of HTML files to be served statically by a web server.",
	Run:   runBuildCommand,
}

func runBuildCommand(cmd *cobra.Command, args []string) {
	// TODO
}
