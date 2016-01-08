/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/codegangsta/cli"
)

// TODO: Add flag handling
var NewCmd = cli.Command{
	Name:            "new",
	Aliases:         []string{"n"},
	Usage:           "hyde new <name> - Creates a new Hyde site",
	SkipFlagParsing: true,
	Action:          doNew,
}

func doNew(ctx *cli.Context) {
	// TODO
}
