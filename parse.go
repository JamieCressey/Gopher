package main

import (
	"github.com/vjeantet/grok"
)

func GrokParse(line, grokFilter string) map[string]string {
	g, _ := grok.New()

	values, _ := g.Parse(grokFilter, line)

	return values
}
