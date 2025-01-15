package stringex

import "strings"

func TrimSpaceToLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}
