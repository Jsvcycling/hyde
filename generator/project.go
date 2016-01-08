/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"os"
	"path"
)

var (
	workingDir string
	subdirs    = []string{
		"pages",
		"public",
		path.Join("public", "javascript"),
		path.Join("public", "styles"),
		path.Join("public", "images"),
	}
)

func init() {
	// Figure out the current working directory
	var err error
	workingDir, err = os.Getwd()

	if err != nil {
		panic(err)
	}
}

func CreateNewProject(targetDir string) error {
	// First make the target directory
	err := os.Mkdir(path.Join(workingDir, targetDir), os.ModeDir)

	if err != nil {
		return err
	}

	// Now make all our subdirectories
	for _, dir := range subdirs {
		err = os.Mkdir(path.Join(workingDir, targetDir, dir), os.ModeDir)

		if err != nil {
			return err
		}
	}

	// TODO: Create our template files

	return nil
}
