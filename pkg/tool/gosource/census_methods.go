package gosource

import "github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"

func CensusMethods(
	c client.Census,
	method string,
) client.Census {
	if method == "" {
		return c
	}

	kept := []client.CensusMethod{}

	for _, m := range c.Methods {
		if m.Name == method {
			kept = append(kept, m)
		}
	}

	c.Methods = kept

	return c
}
