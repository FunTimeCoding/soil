package command_recorder

import "fmt"

func (l *Recorder) BeginCommand(name string) {
	l.Entries = append(l.Entries, fmt.Sprintf("begin %s", name))
}
