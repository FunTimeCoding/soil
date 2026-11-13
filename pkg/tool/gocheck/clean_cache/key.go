package clean_cache

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/string_constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/unclosed_resource"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"strconv"
)

func (c *Cache) Key(u *module_graph.Node) (string, bool) {
	unit, okay := c.handle.UnitKey(u.Path)

	if !okay {
		return "", false
	}

	closure, okay := c.handle.ClosureKey(u.Path, unclosed_resource.Kind())

	if !okay {
		return "", false
	}

	h := sha256.New()
	write := func(s ...string) {
		for _, value := range s {
			h.Write([]byte(value))
			h.Write([]byte{0})
		}
	}
	write(unit, closure, c.faces, strconv.FormatBool(c.comment))
	write(c.ours...)

	for _, path := range string_constant.ListedConstantFiles(u.Directory, c.list) {
		digest, okay := c.constantDigest(path)

		if !okay {
			return "", false
		}

		write(path, digest)
	}

	return hex.EncodeToString(h.Sum(nil)), true
}
