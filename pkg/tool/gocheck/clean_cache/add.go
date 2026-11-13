package clean_cache

import "github.com/funtimecoding/soil/pkg/tool/gocheck/constant"

func (c *Cache) Add(key string) {
	c.handle.Mark(constant.CleanMarks, key)
}
