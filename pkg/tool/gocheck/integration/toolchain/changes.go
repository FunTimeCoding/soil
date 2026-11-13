package toolchain

import (
	"github.com/funtimecoding/soil/pkg/source/types/layer_change"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"slices"
	"strings"
)

func changes(
	r *check_result.Result,
	path string,
) bool {
	return slices.ContainsFunc(
		r.Changes,
		func(c *layer_change.Change) bool {
			return strings.HasSuffix(c.Path, path)
		},
	)
}
