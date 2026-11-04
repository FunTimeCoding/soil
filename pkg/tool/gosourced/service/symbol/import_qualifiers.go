package symbol

import (
	"fmt"
	"go/types"
)

func importQualifiers(
	imports []string,
	index map[string]*types.Package,
) (map[string]string, error) {
	result := map[string]string{}

	for _, path := range imports {
		p, okay := index[path]

		if !okay {
			return nil, fmt.Errorf(
				"import %s is not in the loaded build graph",
				path,
			)
		}

		result[p.Name()] = path
	}

	return result, nil
}
