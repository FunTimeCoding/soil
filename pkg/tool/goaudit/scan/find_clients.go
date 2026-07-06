package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"path/filepath"
)

func findClients(
	v *virtual_file_system.System,
	root string,
	current string,
	repo string,
	configuration *Configuration,
) []*Client {
	var result []*Client

	if v.Has(filepath.Join(current, "client.go")) &&
		!isExcluded(root, current, configuration) {
		c := scanClient(v, root, current, repo)
		result = append(result, c)

		return result
	}

	for _, n := range v.MustReadDirectory(current) {
		if isExcludedDirectory(n) {
			continue
		}

		c := filepath.Join(current, n)

		if !v.DirectoryExists(c) {
			continue
		}

		result = append(
			result,
			findClients(v, root, c, repo, configuration)...,
		)
	}

	return result
}
