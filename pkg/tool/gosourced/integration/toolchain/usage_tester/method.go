package usage_tester

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"

func (o *Tester) Method(
	c *result.Usage,
	name string,
) *result.UsageMethod {
	o.t.Helper()

	for _, m := range c.Methods {
		if m.Name == name {
			return m
		}
	}

	o.t.Fatalf("usage has no method %s", name)

	return nil
}
