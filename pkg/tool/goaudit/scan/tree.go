package scan

import "github.com/funtimecoding/soil/pkg/system/virtual_file_system"

func Tree(root string) *virtual_file_system.System {
	whole := virtual_file_system.From(root)
	result := virtual_file_system.New()

	for _, path := range whole.Files() {
		if isIgnoredPath(path) {
			continue
		}

		f := whole.FileAt(path)
		result.AddFile(path, f.Content, f.ModTime)
	}

	return result
}
