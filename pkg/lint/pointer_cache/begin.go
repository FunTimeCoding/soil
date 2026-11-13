package pointer_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"slices"
)

func (c *Cache) Begin(changes map[string][]string) {
	c.checks++
	c.auditing = c.checks%constant.PointerAuditInterval == 0
	relatives := make(map[string][]string)

	for k, e := range c.entries {
		changed, done := relatives[e.Root]

		if !done {
			changed = relativeChanges(e.Root, changes)
			relatives[e.Root] = changed
		}

		if slices.ContainsFunc(changed, globalChange) ||
			e.Dependencies.Touched(changed) {
			delete(c.entries, k)
		}
	}
}
