package file_identity

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func skip(name string) bool {
	if strings.HasSuffix(name, soil.TestSuffix) {
		return true
	}

	base := strings.TrimSuffix(name, soil.GoExtension)

	if base == "constant" || strings.HasPrefix(base, "constant_") {
		return true
	}

	return false
}
