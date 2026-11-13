package document_cache

import "github.com/funtimecoding/soil/pkg/lint/file_report"

func (c *Cache) Rotate() {
	c.previous = c.current
	c.current = make(map[string]*file_report.Report)
}
