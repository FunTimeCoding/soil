package usage_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"testing"
)

func Method(
	t *testing.T,
	c *result.Usage,
	name string,
) *result.UsageMethod {
	t.Helper()

	for _, m := range c.Methods {
		if m.Name == name {
			return m
		}
	}

	t.Fatalf("usage has no method %s", name)

	return nil
}
