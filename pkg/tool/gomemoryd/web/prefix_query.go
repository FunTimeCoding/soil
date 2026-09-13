package web

import (
	library "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func PrefixQuery(v string) string {
	fields := strings.Fields(v)

	if len(fields) == 0 {
		return v
	}

	last := fields[len(fields)-1]

	if !plainToken(last) {
		return v
	}

	fields[len(fields)-1] = join.Empty(last, constant.PrefixMark)

	return strings.Join(fields, library.Space)
}
