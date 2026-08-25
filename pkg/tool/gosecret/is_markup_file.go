package gosecret

import (
	"path/filepath"
	"strings"
)

func isMarkupFile(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))

	return extension == ".yaml" || extension == ".yml"
}
