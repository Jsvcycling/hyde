/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/spf13/cobra"
)

var genCommand = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"generate"},
	Short:   "Generate a set of HTML files to be served statically by a web server.",
	Run:     runGenCommand,
}

func runGenCommand(cmd *cobra.Command, args []string) {
	// TODO
}
