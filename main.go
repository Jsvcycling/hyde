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
	app.Name = "hyde"
	app.Usage = "a lightweight static site generator"
	app.Version = "0.0.1 ALPHA"
	app.Commands = []cli.Command{
		cmds.BuildCmd,
		cmds.GenCmd,
		cmds.NewCmd,
		cmds.ServeCmd,
	}

	app.Run(os.Args)
}
