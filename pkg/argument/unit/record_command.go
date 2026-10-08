package unit

import "fmt"

func (l *CommandLog) RecordCommand(name string) {
	l.entries = append(l.entries, fmt.Sprintf("record %s", name))
}
