/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"os"

	"github.com/BurntSushi/toml"
)

type ConfigData struct {
	Port  int
	Error error
}

func ParseConfig(filename string) *ConfigData {
	var config *ConfigData
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		return &ConfigData{Error: err}
	}

	if _, err = toml.DecodeReader(file, config); err != nil {
		return &ConfigData{Error: err}
	}

	if config.Port == 0 {
		config.Port = 3000
	}

	return config
}
