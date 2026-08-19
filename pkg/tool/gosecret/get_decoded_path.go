package gosecret

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"path/filepath"
	"strings"
)

func GetDecodedPath(path string) string {
	ext := filepath.Ext(path)
	base := strings.TrimSuffix(path, ext)

	return join.Empty(base, ".decoded.txt")
}
