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
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

var (
	CantOpenFile      = errors.New("Cannot open file")
	CantReadFile      = errors.New("Cannot read file")
	CantParseMetadata = errors.New("Cannot parse metadata")
)

type PageOutput struct {
	Name    string
	Params  map[string]string
	Content string
	Error   error
}

func ParsePage(filename string) *PageOutput {
	file, err := os.Open(filename)
	var parser pageParser

	if err != nil {
		return &PageOutput{Name: filename, Error: CantOpenFile}
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
}

// Creole Parser
type creoleParser struct{}

func (parser creoleParser) fromBuffer(buf io.Reader, output *PageOutput) {
	// TODO: Add Creole support.
}

// HTML Parser
type htmlParser struct{}

func (parser htmlParser) fromBuffer(buf io.Reader, output *PageOutput) {
	output.Params = parseMetadata(buf)

	data, err := ioutil.ReadAll(buf)

	if err != nil {
		// Clear params if we hit an error
		output.Params = nil
		output.Error = CantReadFile
		return
	}

	// Don't use SanitizeReader cause we can't catch read errors
	output.Content = string(bluemonday.UGCPolicy().SanitizeBytes(data))
}

// Textile Parser
type textileParser struct{}

func (parser textileParser) fromBuffer(buf io.Reader, output *PageOutput) {
	// TODO: Add Textile support.
}

// Markdown Parser
type markdownParser struct{}

func (parser markdownParser) fromBuffer(buf io.Reader, output *PageOutput) {
	// TODO: Add Markdown support.
}
