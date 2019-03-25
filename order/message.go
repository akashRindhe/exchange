package order

import "strings"

type Message map[string]string

func (m Message) ToString() string {
	var sb strings.Builder
	for k := range m {
		sb.WriteString(k)
		sb.WriteString("=")
		sb.WriteString(m[k])
		sb.WriteString(";")
	}
	return sb.String()
}
