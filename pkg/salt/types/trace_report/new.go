package trace_report

import "github.com/funtimecoding/soil/pkg/salt/types/top_entry"

func New(
	target string,
	entries []*top_entry.Entry,
	states []*Node,
	files []*Node,
	missing []string,
) *Report {
	return &Report{
		Target: target,
		Summary: &Summary{
			States:         len(states),
			Files:          len(files),
			OrphanedStates: orphaned(states),
			OrphanedFiles:  orphaned(files),
		},
		Entries: entries,
		States:  states,
		Files:   files,
		Missing: missing,
	}
}
