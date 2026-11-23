package state_tree

import "slices"

func (t *Tree) reachedStates(
	roots []string,
) (map[string]bool, map[string]bool) {
	reached := map[string]bool{}
	missing := map[string]bool{}
	queue := slices.Clone(roots)

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if reached[s] {
			continue
		}

		if _, exists := t.states[s]; !exists {
			missing[s] = true

			continue
		}

		reached[s] = true
		queue = append(queue, t.includes[s]...)
	}

	return reached, missing
}
