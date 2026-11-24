package installed

import "github.com/funtimecoding/soil/pkg/go_mod"

func requires(directory string) map[string]bool {
	result := make(map[string]bool)

	for _, r := range go_mod.ReadDirectory(directory).Require {
		result[r.Mod.Path] = true
	}

	return result
}
