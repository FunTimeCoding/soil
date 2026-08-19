package gosecret

import (
	"path/filepath"
	"strings"
)

func isYAMLFile(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))

	return extension == ".yaml" || extension == ".yml"
}
