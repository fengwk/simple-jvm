package common

import "strings"

const INDENT2 = "  "
const INDENT4 = "    "

func Indent(s string, indent string) string {
	sb := strings.Builder{}
	lines := strings.Split(s, "\n")
	for i, s := range lines {
		sb.WriteString(indent)
		sb.WriteString(s)
		if i < len(lines)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
