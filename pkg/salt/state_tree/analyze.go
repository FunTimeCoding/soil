package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/types/analysis"
	"strings"
)

func (t *Tree) Analyze() *analysis.Analysis {
	reached, possible, missing := t.walk(t.entriesWhere(isLive))
	live := map[string]bool{}
	unreached := map[string]bool{}

	for s := range t.states {
		switch {
		case reached[s] || possible[s]:
			live[s] = true
		default:
			unreached[s] = true
		}
	}

	used := t.usedFiles(live)
	unusedFiles := map[string]bool{}

	for f := range t.files {
		if !used[f] && !strings.HasSuffix(f, constant.StateExtension) {
			unusedFiles[f] = true
		}
	}

	result := analysis.New()
	result.Reached = sortedKeys(reached)
	result.Possible = sortedKeys(possible)
	result.Unreached = sortedKeys(unreached)
	result.Missing = sortedKeys(missing)
	result.UnreachedFiles = sortedKeys(unusedFiles)
	result.DeadEnvironments = sortedKeys(t.dead)

	return result
}
