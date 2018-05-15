package ninja

import "strings"

func ParseSpecialText(str string) string {
	str = strings.Replace(str, "|", "\n", -1)
	s := strings.Index(str, "{")
	if s == -1 {
		return str
	}
	s += len("{")
	e := strings.Index(str, "}")
	return str[s:e]
}
