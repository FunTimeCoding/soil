package clean_cache

import (
	"crypto/sha256"
	"encoding/hex"
)

func (c *Cache) constantDigest(path string) (string, bool) {
	if cached, found := c.constants[path]; found {
		return cached, true
	}

	content, e := c.layer.Read(path)

	if e != nil {
		return "", false
	}

	sum := sha256.Sum256(content)
	result := hex.EncodeToString(sum[:])
	c.constants[path] = result

	return result, true
}
