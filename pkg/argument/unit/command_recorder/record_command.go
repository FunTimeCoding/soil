package command_recorder

import "fmt"

func (l *Recorder) RecordCommand(name string) {
	l.Entries = append(l.Entries, fmt.Sprintf("record %s", name))
}
