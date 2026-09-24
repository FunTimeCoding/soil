package mutation

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Load(path string) *Report {
	var result Report
	notation.MustDecodeBytes(system.ReadBytesUnsafe(path), &result, false)

	return &result
}
