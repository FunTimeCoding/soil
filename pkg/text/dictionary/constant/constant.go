package constant

import "github.com/funtimecoding/soil/pkg/project"

const File = "dictionary.dic"

var (
	Skip   = map[string]bool{".git": true, "tmp": true}
	Prefix = map[string]bool{
		"Containerfile.": true,
		"Dockerfile.":    true,
	}
	Extension = map[string]bool{
		".conf":      true,
		".envrc":     true,
		".gitignore": true,
		".go":        true,
		".json":      true,
		".md":        true,
		".rego":      true,
		".sh":        true,
		".txt":       true,
		".xml":       true,
		".yaml":      true,
		".yml":       true,
	}
	NoExtension = map[string]bool{
		project.ContainerFile: true,
		project.DockerFile:    true,
	}
)
