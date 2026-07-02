package function

import "slices"

func validTool(name string) bool {
	var valid []string

	for _, v := range functions {
		valid = append(valid, v.Name)
	}

	return slices.Contains(valid, name)
}
