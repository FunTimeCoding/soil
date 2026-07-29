package virtual_file_system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"sort"
	"strings"
)

func (s *System) ReadDirectory(path string) ([]string, error) {
	if !s.DirectoryExists(path) {
		return nil, constant.ErrorDirectoryNotFound
	}

	prefix := fmt.Sprintf("%s/", path)
	seen := make(map[string]bool)

	for key := range s.files {
		if !strings.HasPrefix(key, prefix) {
			continue
		}

		remainder := strings.TrimPrefix(key, prefix)
		name, _, _ := strings.Cut(remainder, "/")
		seen[name] = true
	}

	result := make([]string, 0, len(seen))

	for name := range seen {
		result = append(result, name)
	}

	sort.Strings(result)

	return result, nil
}
