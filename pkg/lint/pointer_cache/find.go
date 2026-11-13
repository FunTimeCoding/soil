package pointer_cache

import "github.com/funtimecoding/soil/pkg/lint/types/pointer_entry"

func (c *Cache) Find(
	root string,
	document string,
	content string,
) (*pointer_entry.Entry, bool) {
	e, found := c.entries[entryKey(root, document)]

	if !found || e.Key != c.contentKey(content) {
		return nil, false
	}

	return e, true
}
