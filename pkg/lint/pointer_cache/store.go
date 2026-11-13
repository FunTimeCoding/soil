package pointer_cache

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_dependency"
	"github.com/funtimecoding/soil/pkg/lint/types/pointer_entry"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
)

func (c *Cache) Store(
	root string,
	document string,
	content string,
	report *file_report.Report,
	spans []*unchecked.Unchecked,
	dependencies *pointer_dependency.Set,
) {
	k := entryKey(root, document)

	if dependencies.Outside {
		delete(c.entries, k)

		return
	}

	c.entries[k] = pointer_entry.New(
		root,
		c.contentKey(content),
		report.Clone(),
		spans,
		dependencies,
	)
}
