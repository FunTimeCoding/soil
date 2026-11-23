package state_tree

import (
	"regexp"
	"slices"
)

func (t *Tree) possibleStates(
	reached map[string]bool,
	rootPatterns []*regexp.Regexp,
) map[string]bool {
	result := map[string]bool{}
	patterns := slices.Clone(rootPatterns)

	for s := range reached {
		patterns = append(patterns, t.statePattern[s]...)
	}

	queue := t.matching(patterns)

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if reached[s] || result[s] {
			continue
		}

		if _, exists := t.states[s]; !exists {
			continue
		}

		result[s] = true
		queue = append(queue, t.includes[s]...)
		queue = append(
			queue,
			t.matching(append([]*regexp.Regexp{}, t.statePattern[s]...))...,
		)
	}

	return result
}
