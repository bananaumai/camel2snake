package converter

import (
	"regexp"
	"strings"
)

type Converter struct {
	re1 *regexp.Regexp
	re2 *regexp.Regexp
}

func NewConverter() *Converter {
	return &Converter{
		regexp.MustCompile("(.)([A-Z][a-z]+)"),
		regexp.MustCompile("([a-z0-9])([A-Z])"),
	}
}

func (c *Converter) Convert(str string) string {
	str = c.re1.ReplaceAllString(str, "${1}_${2}")
	str = c.re2.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(str)
}
