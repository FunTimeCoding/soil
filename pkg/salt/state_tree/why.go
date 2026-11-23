package state_tree

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/salt/types/chain"
	"github.com/funtimecoding/soil/pkg/salt/types/why_report"
)

func (t *Tree) Why(state string) (*why_report.Report, error) {
	if _, exists := t.states[state]; !exists {
		return nil, not_found.New("state", state)
	}

	reached, possible, _ := t.walk(t.entriesWhere(isLive))
	status := constant.StatusUnreached

	switch {
	case reached[state]:
		status = constant.StatusReached
	case possible[state]:
		status = constant.StatusPossible
	}

	var chains []*chain.Chain

	for _, e := range t.entries {
		if c := t.shortestChain(e.State, state); c != nil {
			chains = append(chains, chain.New(e, c))
		}
	}

	return why_report.New(state, status, chains), nil
}
