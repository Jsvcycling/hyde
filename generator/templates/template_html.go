/*
 * Copyright (c) 2016 Josh Vega
 * See LICENSE for license details.
 */
package templates

var TemplateHTML string = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta name="author" content="{{author}}">
        <meta http-equiv="last-modified" content="{{date}}">
        <meta name="description" content="{{description}}">
        <meta name="keywords" content="{{keywords}}">
        <meta name="generator" content="Hyde by Jsvcycling (github.com/jsvcycling/hyde)">

        <title>{{title}}</title>
    </head>
    <body>
        {{content}}
    </body>
</html>`