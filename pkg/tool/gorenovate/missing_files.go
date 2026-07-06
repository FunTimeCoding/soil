package gorenovate

import "github.com/funtimecoding/soil/pkg/system"

func missingFiles(c *Configuration) []string {
	var result []string

	for _, r := range c.PackageRules {
		for _, f := range r.MatchFiles {
			if !system.FileExists(f) {
				result = append(result, f)
			}
		}
	}

	return result
}
