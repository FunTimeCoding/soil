package check

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lint/pointer_cache"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func resolvedCache(r face.Reporter) *pointer_cache.Cache {
	result := pointer_cache.New()
	result.Mismatch = func(
		root string,
		document string,
		cached string,
		fresh string,
	) {
		r.CaptureWithContext(
			unexpected.Format(
				"pointer cache missed a dependency: %s",
				document,
			),
			constant.PointerCacheContext,
			map[string]any{"root": root, "cached": cached, "fresh": fresh},
		)
	}

	return result
}
