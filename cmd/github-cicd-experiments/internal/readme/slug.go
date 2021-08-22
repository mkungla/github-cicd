package readme

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	// NamespaceMustCompile against following expression.
	NamespaceMustCompile = "^[a-zA-Z][a-zA-Z0-9_-]*[a-zA-Z0-9]$"
)

var (
	rExps = []replacement{ //nolint:gochecknoglobals
		{re: regexp.MustCompile(`[\xC0-\xC6]`), ch: "A"},
		{re: regexp.MustCompile(`[\xE0-\xE6]`), ch: "a"},
		{re: regexp.MustCompile(`[\xC8-\xCB]`), ch: "E"},
		{re: regexp.MustCompile(`[\xE8-\xEB]`), ch: "e"},
		{re: regexp.MustCompile(`[\xCC-\xCF]`), ch: "I"},
		{re: regexp.MustCompile(`[\xEC-\xEF]`), ch: "i"},
		{re: regexp.MustCompile(`[\xD2-\xD6]`), ch: "O"},
		{re: regexp.MustCompile(`[\xF2-\xF6]`), ch: "o"},
		{re: regexp.MustCompile(`[\xD9-\xDC]`), ch: "U"},
		{re: regexp.MustCompile(`[\xF9-\xFC]`), ch: "u"},
		{re: regexp.MustCompile(`[\xC7-\xE7]`), ch: "c"},
		{re: regexp.MustCompile(`[\xD1]`), ch: "N"},
		{re: regexp.MustCompile(`[\xF1]`), ch: "n"},
	}
	spacereg       = regexp.MustCompile(`\s+`)
	noncharreg     = regexp.MustCompile(`[^A-Za-z0-9-]`)
	minusrepeatreg = regexp.MustCompile(`\-{2,}`)
	alnum          = &unicode.RangeTable{ //nolint:gochecknoglobals
		R16: []unicode.Range16{
			{'0', '9', 1},
			{'A', 'Z', 1},
			{'a', 'z', 1},
		},
	}
)

// Replacement structure.
type replacement struct {
	re *regexp.Regexp
	ch string
}

// SlugOf returns markdown slug representation of given string
func SlugOf(str string) string {
	for _, r := range rExps {
		str = r.re.ReplaceAllString(str, r.ch)
	}

	str = strings.ToLower(str)
	str = spacereg.ReplaceAllString(str, "-")
	str = noncharreg.ReplaceAllString(str, "")
	str = minusrepeatreg.ReplaceAllString(str, "-")

	return str
}
