/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package cmds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/codegangsta/cli"

	"github.com/jsvcycling/hyde/helpers"
	"github.com/jsvcycling/hyde/parser"
)

var (
	workingDir string
)

func init() {
	// Figure out the current working directory
	var err error
	workingDir, err = os.Getwd()

	if err != nil {
		panic(err)
	}
}

// TODO: Add flag handling
var BuildCmd = cli.Command{
	Name:            "build",
	Aliases:         []string{"b"},
	Usage:           "build and compile the site",
	SkipFlagParsing: true,
	Action:          doBuild,
}

func doBuild(ctx *cli.Context) {
	if _, err := os.Stat("hyde.toml"); os.IsNotExist(err) {
		fmt.Println("Hyde project not found")
		return
	}

	// Make sure the pages directory exists
	if _, err := os.Stat("pages"); os.IsNotExist(err) {
		fmt.Println("Invalid Hyde project")
		return
	}

	// Make sure the templates directory exists
	if _, err := os.Stat("templates"); os.IsNotExist(err) {
		fmt.Println("Invalid Hyde project")
		return
	}

	config := parser.ParseConfig("hyde.toml")

	if config.Error != nil {
		fmt.Println(config.Error.Error())
		return
	}

	// If the target directory exists, delete it and recreate it otherwise just
	// create it.
	if _, err := os.Stat(config.TargetDir); !os.IsNotExist(err) {
		if err := os.RemoveAll(config.TargetDir); err != nil {
			fmt.Println(err.Error())
		}
	}

	if err := os.Mkdir(config.TargetDir, os.ModeDir); err != nil {
		fmt.Println(err.Error())
	}

	// Get all the pages in the pages directory
	pages, err := ioutil.ReadDir(path.Join(workingDir, "pages"))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, page := range pages {
		pageData := parser.ParsePage("pages", page.Name())

		if pageData.Error != nil {
			fmt.Println(pageData.Error.Error())
			return
		}

		if err := pageData.GeneratePage(config.TargetDir, config.Minify); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Copy the contents of the public directory directly to the output.
	helpers.CopyDir("public", path.Join(config.TargetDir, "public"))
}
