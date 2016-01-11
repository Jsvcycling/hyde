/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"bufio"
	"errors"
	"io"

	"github.com/BurntSushi/toml"
)

const (
	METADATA_START = "+++"
	METADATA_END   = "+++"
)

var (
	ErrorMissingTitle       = errors.New("Missing title")
	ErrorMissingAuthor      = errors.New("Missing author")
	ErrorMissingDescription = errors.New("Missing description")
	ErrorMissingTemplate    = errors.New("Missing template")
	ErrorNoMetadata         = errors.New("Couldn't find metadata")
)

// TODO: Make certain fields optional
type PageMetadata struct {
	Title       string `toml:"title"`
	Author      string
	Description string
	Template    string
}

func parseMetadata(buf io.Reader, output *PageMetadata) error {
	var config PageMetadata
	reader := bufio.NewReader(buf)

	// Should we use ReadBytes instead?
	line, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	if line[:len(line)-1] == METADATA_START {
		var data string

		// Read each metadata line until the metadata section ends
		for {
			if str, err := reader.ReadString('\n'); err == nil && str[:len(str)-1] != METADATA_END {
				data += str
			} else if err != nil {
				return err
			} else {
				break
			}
		}

		_, err := toml.Decode(data, &config)

		if err != nil {
			return err
		}
	} else {
		return ErrorNoMetadata
	}

	if config.Title == "" {
		return ErrorMissingTitle
	} else if config.Author == "" {
		return ErrorMissingAuthor
	} else if config.Description == "" {
		return ErrorMissingDescription
	} else if config.Template == "" {
		return ErrorMissingTemplate
	}

	output = &config
	return nil
}
