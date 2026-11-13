package pointer_cache

import (
	"fmt"
	"hash/maphash"
)

func (c *Cache) contentKey(content string) string {
	return fmt.Sprintf(
		"%d\x00%x",
		len(content),
		maphash.String(c.seed, content),
	)
}
