/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"errors"
	"io"
	"os"
	"strings"
)

var (
	CantReadFile      = errors.New("Cannot open file")
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
		return &PageOutput{Name: filename, Error: CantReadFile}
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
	// TODO: Do we really need to do anything?
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