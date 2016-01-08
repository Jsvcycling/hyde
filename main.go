/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package main

import (
	"os"

	"github.com/codegangsta/cli"

	"github.com/jsvcycling/hyde/cmds"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		cmds.BuildCmd,
		cmds.GenCmd,
		cmds.NewCmd,
		cmds.ServeCmd,
	}

	app.Run(os.Args)
}
