package entry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
)

func (e *Entry) Location(root string) string {
	return fmt.Sprintf(
		"%s:%d",
		system.RelativePath(root, e.Function.File),
		e.Function.Line,
	)
}
