package pointer_cache

func (c *Cache) Validate(
	root string,
	global string,
) {
	if c.globals[root] == global {
		return
	}

	for k, e := range c.entries {
		if e.Root == root {
			delete(c.entries, k)
		}
	}

	c.globals[root] = global
}
