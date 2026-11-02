package toolchain

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"sort"
)

func remainingShapes(c *result.Constructor) []string {
	var shapes []string

	for _, g := range c.Remaining {
		shapes = append(shapes, g.Shape)
	}

	sort.Strings(shapes)

	return shapes
}
