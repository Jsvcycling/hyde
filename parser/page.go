/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/m4tty/cajun"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var (
	ErrorNotImplemented = errors.New("Not implemented")
	ErrorBadWrite       = errors.New("Unable to write to output file")
)

type PageOutput struct {
	Name     string
	Metadata PageMetadata
	Content  string
	Error    error
}

func ParsePage(filename string) *PageOutput {
	file, err := os.Open(filename)
	var parser pageParser

	if err != nil {
		return &PageOutput{Name: filename, Error: err}
	}

	switch guessTypeByExt(filename) {
	case "asciidoc":
		parser = asciiDocParser{}
	case "creole":
		parser = creoleParser{}
	case "markdown":
		parser = markdownParser{}
	case "textile":
		parser = textileParser{}
	default:
		parser = htmlParser{}
	}

	var output PageOutput

	output.Name = filename[:strings.LastIndex(filename, ".")]

	parser.fromBuffer(file, &output)

	return &output
}

// TODO: Add support for minification
func (page *PageOutput) GeneratePage(targetDir string, minify bool) error {
	path, err := filepath.Abs(targetDir + "/" + page.Name + ".html")

	if err != nil {
		return err
	}

	outFile, err := os.Create(path)
	defer outFile.Close()

	if err != nil {
		return err
	}

	tmplPath, err := filepath.Abs("templates/" + page.Metadata.Template + ".html")

	if err != nil {
		return err
	}

	tmplFile, err := os.Open(tmplPath)
	defer tmplFile.Close()

	if err != nil {
		return err
	}

	tmplBytes, err := ioutil.ReadAll(tmplFile)

	if err != nil {
		return err
	}

	tmplData := string(tmplBytes)

	// Embed page information into the template.
	tmplData = strings.Replace(tmplData, "{{title}}", page.Metadata.Title, -1)
	tmplData = strings.Replace(tmplData, "{{author}}", page.Metadata.Author, -1)
	tmplData = strings.Replace(tmplData, "{{description}}", page.Metadata.Description, -1)
	tmplData = strings.Replace(tmplData, "{{date}}", page.Metadata.Date.String(), -1)
	tmplData = strings.Replace(tmplData, "{{content}}", page.Content, -1)

	// Resanitize our HTML (just in case)
	tmplData = bluemonday.UGCPolicy().Sanitize(tmplData)

	chars, err := outFile.WriteString(tmplData)

	if err != nil {
		return err
	} else if chars != len(tmplData) {
		// Make sure we write the entire template
		return ErrorBadWrite
	}

	return nil
}

func guessTypeByExt(filename string) string {
	idx := strings.LastIndex(filename, ".")

	switch strings.ToLower(filename[idx+1:]) {
	case "asciidoc", "ascii":
		return "asciidoc"
	case "creole":
		return "creole"
	case "md", "markdown":
		return "markdown"
	case "textile":
		return "textile"
	default:
		return "html"
	}
}

type pageParser interface {
	fromBuffer(buf io.Reader, output *PageOutput)
}

// AsciiDoc Parser
type asciiDocParser struct{}

func (parser asciiDocParser) fromBuffer(buf io.Reader, output *PageOutput) {
	// TODO: Add AsciiDoc support.
	output.Error = ErrorNotImplemented
}

// Creole Parser
type creoleParser struct{}

func (parser creoleParser) fromBuffer(buf io.Reader, output *PageOutput) {
	if err := parseMetadata(buf, &output.Metadata); err != nil {
		output.Error = err
		return
	}

	data, err := ioutil.ReadAll(buf)

	if err != nil {
		output.Error = err
		return
	}

	content, err := cajun.Transform(string(data))

	if err != nil {
		output.Error = err
		return
	}

	output.Content = bluemonday.UGCPolicy().Sanitize(content)
}

// HTML Parser
type htmlParser struct{}

func (parser htmlParser) fromBuffer(buf io.Reader, output *PageOutput) {
	var data []byte

	err := parseMetadata(buf, &output.Metadata)

	if err != nil {
		output.Error = err
		return
	}

	// Does this read from the cursors current location?
	data, err = ioutil.ReadAll(buf)

	if err != nil {
		output.Error = err
		return
	}

	// Don't use SanitizeReader cause we can't catch read errors
	output.Content = string(bluemonday.UGCPolicy().SanitizeBytes(data))
}

// Textile Parser
type textileParser struct{}

func (parser textileParser) fromBuffer(buf io.Reader, output *PageOutput) {
	// TODO: Add Textile support.
	output.Error = ErrorNotImplemented
}

// Markdown Parser
type markdownParser struct{}

func (parser markdownParser) fromBuffer(buf io.Reader, output *PageOutput) {
	var data []byte

	err := parseMetadata(buf, &output.Metadata)

	if err != nil {
		output.Error = err
		return
	}

	// Does this read from the cursors current location?
	data, err = ioutil.ReadAll(buf)

	if err != nil {
		output.Error = err
		return
	}

	output.Content = string(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon(data)))
}
