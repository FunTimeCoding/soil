package state_tree

import "regexp"

func (t *Tree) matching(patterns []*regexp.Regexp) []string {
	var result []string

	for name := range t.states {
		for _, p := range patterns {
			if p.MatchString(name) {
				result = append(result, name)

				break
			}
		}
	}

	return result
}
