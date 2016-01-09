/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"fmt"

	"github.com/codegangsta/cli"

	"github.com/jsvcycling/hyde/generator"
)

// TODO: Add flag handling
var NewCmd = cli.Command{
	Name:            "new",
	Aliases:         []string{"n"},
	Usage:           "creates a new Hyde site",
	ArgsUsage:       "[name]",
	SkipFlagParsing: true,
	Action:          doNew,
}

func doNew(ctx *cli.Context) {
	if !ctx.Args().Present() {
		// fmt.Println("ERROR: Missing required argument")
		cli.ShowCommandHelp(ctx, "new")
		return
	}

	if err := generator.CreateNewProject(ctx.Args().First()); err != nil {
		fmt.Println(err.Error())
		return
	}
}
