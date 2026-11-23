package state_tree

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/types/top_entry"
	"github.com/funtimecoding/soil/pkg/salt/types/trace_report"
	"strings"
)

func (t *Tree) Trace(target string) (*trace_report.Report, error) {
	selected := t.entriesWhere(
		func(e *top_entry.Entry) bool { return e.Target == target },
	)

	if len(selected) == 0 {
		return nil, not_found.New("top target", target)
	}

	closure, missing := t.liveStates(selected)
	remaining, _ := t.liveStates(
		t.entriesWhere(
			func(e *top_entry.Entry) bool {
				return e.Live && e.Target != target
			},
		),
	)
	var states []*trace_report.Node

	for _, s := range sortedKeys(closure) {
		states = append(states, trace_report.NewNode(s, !remaining[s]))
	}

	baseline := t.usedFiles(map[string]bool{})
	kept := t.usedFiles(remaining)
	var files []*trace_report.Node

	for _, f := range sortedKeys(t.usedFiles(closure)) {
		if baseline[f] || strings.HasSuffix(f, constant.StateExtension) {
			continue
		}

		files = append(files, trace_report.NewNode(f, !kept[f]))
	}

	return trace_report.New(
		target,
		selected,
		states,
		files,
		sortedKeys(missing),
	), nil
}
