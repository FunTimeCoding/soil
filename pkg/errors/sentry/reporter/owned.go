package reporter

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Reporter) owned(function string) bool {
	return r.module != "" &&
		strings.HasPrefix(function, join.Empty(r.module, constant.Slash)) &&
		!strings.Contains(packagePath(function), "/generated/")
}
