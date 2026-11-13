package clean_cache

import "github.com/funtimecoding/soil/pkg/tool/gocheck/constant"

func (c *Cache) Holds(key string) bool {
	return c.handle.Marked(constant.CleanMarks, key)
}
