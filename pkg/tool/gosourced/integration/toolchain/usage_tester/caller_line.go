package usage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func CallerLine(c *result.UsageCaller) string {
	return fmt.Sprintf("%s %s %s %d", c.Kind, c.Package, c.Function, c.Count)
}
