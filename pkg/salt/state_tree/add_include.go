package state_tree

import "strings"

func (t *Tree) addInclude(
	owner string,
	target string,
) {
	if strings.Contains(target, "{{") {
		t.statePattern[owner] = append(
			t.statePattern[owner],
			dynamicPattern(target),
		)

		return
	}

	t.includes[owner] = append(t.includes[owner], target)
}
