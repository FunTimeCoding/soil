package unit

import "fmt"

func (l *CommandLog) BeginCommand(name string) {
	l.entries = append(l.entries, fmt.Sprintf("begin %s", name))
}
