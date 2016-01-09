/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package generator

import (
	"errors"
	"os"
	"path"

	"github.com/jsvcycling/hyde/generator/templates"
)

var (
	workingDir string
	subdirs    = []string{
		"pages",
		"public",
		"templates",
		path.Join("public", "javascript"),
		path.Join("public", "styles"),
		path.Join("public", "images"),
	}
	files = []file{
		file{path.Join("pages", "index.md"), templates.IndexMD},
		file{path.Join("templates", "template.html"), templates.TemplateHTML},
		file{path.Join("public", "styles", "main.css"), templates.MainCSS},
		file{path.Join("hyde.toml"), templates.HydeTOML},
	}
)

type file struct {
	filename string
	template string
}

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

	// Finally, create our files
	for _, filedata := range files {
		file, err := os.Create(path.Join(workingDir, targetDir, filedata.filename))

		if err != nil {
			return err
		}

		count, err := file.WriteString(filedata.template)

		if err != nil {
			return err
		} else if count != len(filedata.template) {
			return errors.New("Error writing")
		}
	}

	return nil
}
