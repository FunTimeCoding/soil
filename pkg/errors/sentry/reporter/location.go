package reporter

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func (r *Reporter) location(f runtime.Frame) string {
	return fmt.Sprintf(
		"%s:%d",
		filepath.Join(r.label(f.Function), filepath.Base(f.File)),
		f.Line,
	)
}
