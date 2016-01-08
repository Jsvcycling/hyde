/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/codegangsta/cli"
)

// TODO: Add flag handling
var GenCmd = cli.Command{
	Name:            "generate",
	Aliases:         []string{"g"},
	Usage:           "generate a new resource for the current site",
	SkipFlagParsing: true,
	Action:          doGen,
}

func doGen(ctx *cli.Context) {
	// TODO
}
