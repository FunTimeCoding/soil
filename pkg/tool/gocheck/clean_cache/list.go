package clean_cache

import "github.com/funtimecoding/soil/pkg/lint/analyzer/string_constant"

func (c *Cache) list(directory string) []string {
	if cached, found := c.listed[directory]; found {
		return cached
	}

	result := string_constant.DirectorySources(directory)
	c.listed[directory] = result

	return result
}
