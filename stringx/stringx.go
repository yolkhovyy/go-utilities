package stringx

import (
	"strings"
	"unicode"
)

func Capitalize(in string) string {
	if len(in) == 0 {
		return in
	}

	runes := []rune(in)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}

func TrimSpaceToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
