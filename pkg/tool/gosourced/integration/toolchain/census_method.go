package toolchain

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"testing"
)

func censusMethod(
	t *testing.T,
	c *result.Census,
	name string,
) *result.CensusMethod {
	t.Helper()

	for _, m := range c.Methods {
		if m.Name == name {
			return m
		}
	}

	t.Fatalf("census has no method %s", name)

	return nil
}
