package converter

import (
	"regexp"
	"strings"
)

var re1 = regexp.MustCompile("(.)([A-Z][a-z]+)")
var re2 = regexp.MustCompile("([a-z0-9])([A-Z])")

func Convert(str string) string {
	str = re1.ReplaceAllString(str, "${1}_${2}")
	str = re2.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(str)
}
