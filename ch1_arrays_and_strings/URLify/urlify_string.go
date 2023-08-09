package URLify

import "strings"

func ReplaceSpace(s string) string {
	s = strings.TrimSpace(s)
	return strings.Replace(s, " ", "%20", -1)
}
