package gosource

import "github.com/funtimecoding/soil/pkg/tool/gosourced/generated/client"

func UsageMethods(
	u client.Usage,
	method string,
) client.Usage {
	if method == "" {
		return u
	}

	kept := []client.UsageMethod{}

	for _, m := range u.Methods {
		if m.Name == method {
			kept = append(kept, m)
		}
	}

	u.Methods = kept

	return u
}
