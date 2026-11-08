package toolchain

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func callerLine(c *result.CensusCaller) string {
	return fmt.Sprintf("%s %s %s %d", c.Kind, c.Package, c.Function, c.Count)
}
