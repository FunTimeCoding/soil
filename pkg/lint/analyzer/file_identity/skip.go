package file_identity

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"strings"
)

func skip(name string) bool {
	if strings.HasSuffix(name, constant.TestSuffix) {
		return true
	}

	base := strings.TrimSuffix(name, constant.GoExtension)

	if base == "constant" || strings.HasPrefix(base, "constant_") {
		return true
	}

	return false
}
