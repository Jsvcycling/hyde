/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"github.com/codegangsta/cli"
)

// TODO: Add flag handling
var ServeCmd = cli.Command{
	Name:            "serve",
	Aliases:         []string{"s"},
	Usage:           "hyde serve -- Serve the current site on a development server",
	SkipFlagParsing: true,
	Action:          doServe,
}

func doServe(ctx *cli.Context) {
	// TODO
}
