package policy

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"strings"
)

func Encode(v []*Policy) string {
	if len(v) == 0 {
		return constant.ClearValue
	}

	var fields []string

	for _, p := range v {
		fields = append(fields, p.KeyName, p.Domain, join.Comma(p.Types))
	}

	return strings.Join(fields, constant.RowSeparator)
}
