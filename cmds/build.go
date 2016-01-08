/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/codegangsta/cli"
)

// TODO: Add flag handling
var BuildCmd = cli.Command{
	Name:            "build",
	Aliases:         []string{"b"},
	Usage:           "hyde build -- Build and compile the site",
	SkipFlagParsing: true,
	Action:          doBuild,
}

func doBuild(ctx *cli.Context) {
	// TODO
}
