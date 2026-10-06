package module_graph

import "path/filepath"

func inDirectory(
	directory string,
	names []string,
) []string {
	result := make([]string, 0, len(names))

	for _, n := range names {
		result = append(result, filepath.Join(directory, n))
	}

	return result
}
