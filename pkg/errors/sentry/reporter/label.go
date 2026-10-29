package reporter

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func (r *Reporter) label(function string) string {
	p := packagePath(function)

	if r.module != "" {
		p = strings.TrimPrefix(p, join.Empty(r.module, constant.Slash))
	}

	return p
}
