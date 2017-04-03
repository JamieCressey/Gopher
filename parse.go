package main

import (
	"github.com/vjeantet/grok"
)

func GrokParse(line, grokFilter string) map[string]interface{} {
	g, _ := grok.New()

	values, _ := g.ParseTyped(grokFilter, line)

	return values
}
