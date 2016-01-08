/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"bufio"
	"errors"
	"io"
	"time"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

const (
	TOML_METADATA_START = "+++"
	TOML_METADATA_END   = "+++"

	YAML_METADATA_START = "---"
	YAML_METADATA_END   = "---"
)

var (
	ErrorMissingTitle       = errors.New("Missing title")
	ErrorMissingAuthor      = errors.New("Missing author")
	ErrorMissingDescription = errors.New("Missing description")
	ErrorMissingDate        = errors.New("Missing date")
	ErrorMissingTemplate    = errors.New("Missing template")
)

// TODO: Make certain fields optional
type PageMetadata struct {
	Title       string
	Author      string
	Description string
	Date        time.Time
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

	if line == TOML_METADATA_START {
		var data string

		// Read each metadata line until the metadata section ends
		for {
			if str, err := reader.ReadString('\n'); err == nil && str != TOML_METADATA_END {
				data += str
				data += "\n"
			} else if err != nil {
				return err
			} else {
				break
			}
		}

		if _, err := toml.Decode(data, &config); err != nil {
			return err
		}
	} else if line == YAML_METADATA_START {
		var data string

		// Read each metadata line until the metadata section ends
		for {
			if str, err := reader.ReadString('\n'); err == nil && str != YAML_METADATA_END {
				data += str
				data += "\n"
			} else if err != nil {
				return err
			} else {
				break
			}
		}

		if err := yaml.Unmarshal([]byte(data), &config); err != nil {
			return err
		}
	}

	if config.Title == "" {
		return ErrorMissingTitle
	} else if config.Author == "" {
		return ErrorMissingAuthor
	} else if config.Description == "" {
		return ErrorMissingDescription
	} else if config.Date.IsZero() {
		return ErrorMissingDate
	} else if config.Template == "" {
		return ErrorMissingTemplate
	}

	output = &config
	return nil
}
