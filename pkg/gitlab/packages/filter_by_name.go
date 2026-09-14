package packages

import "gitlab.com/gitlab-org/api/client-go/v3"

func FilterByName(
	v []*gitlab.Package,
	name string,
) []*gitlab.Package {
	var result []*gitlab.Package

	for _, e := range v {
		if e.Name == name {
			result = append(result, e)
		}
	}

	return result
}
