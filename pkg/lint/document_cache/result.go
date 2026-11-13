package document_cache

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"hash/maphash"
)

func (c *Cache) Result(
	step string,
	path string,
	content string,
	run func() *file_report.Report,
) *file_report.Report {
	key := fmt.Sprintf(
		"%s\x00%s\x00%d\x00%x",
		step,
		path,
		len(content),
		maphash.String(c.seed, content),
	)

	if cached, found := c.current[key]; found {
		return cached.Clone()
	}

	if cached, found := c.previous[key]; found {
		c.current[key] = cached

		return cached.Clone()
	}

	result := run()
	c.current[key] = result.Clone()

	return result
}
