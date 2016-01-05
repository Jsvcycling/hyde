/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package parser

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

// TODO: Write custom Markdown parser based on CommonMark spec.

func ParseBytes(data string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(data))
	return bluemonday.UGCPolicy().Sanitize(string(unsafe))
}
