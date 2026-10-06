package unit

import "fmt"

func (l *commandLog) BeginCommand(name string) {
	l.entries = append(l.entries, fmt.Sprintf("begin %s", name))
}
