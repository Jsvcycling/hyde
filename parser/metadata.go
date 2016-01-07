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
	TOML_METADATA_START = "+++"
	TOML_METADATA_END   = "+++"

	YAML_METADATA_START = "---"
	YAML_METADATA_END   = "---"
)

var (
	ErrorMissingTitle       = errors.New("Missing title")
	ErrorMissingDescription = errors.New("Missing description")
	ErrorMissingDate        = errors.New("Missing date")
)

type PageMetadata struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	Date        string `toml:"date"`
}

func parseMetadata(buf io.Reader) (*PageMetadata, error) {
	var config PageMetadata
	reader := bufio.NewReader(buf)

	// Should we use ReadBytes instead?
	line, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	if line == TOML_METADATA_START {
		var data string

		// Read each metadata line until the metadata section ends
		for {
			if str, err := reader.ReadString('\n'); err == nil && str != TOML_METADATA_END {
				data += str
				data += "\n"
			} else {
				break
			}
		}

		if _, err := toml.Decode(data, &config); err != nil {
			return nil, err
		}
	} else if line == YAML_METADATA_START {
		// TODO: add YAML support
	}

	if config.Title == "" {
		return nil, ErrorMissingTitle
	} else if config.Description == "" {
		return nil, ErrorMissingDescription
	} else if config.Date == "" {
		return nil, ErrorMissingDate
	}

	return &config, nil
}
