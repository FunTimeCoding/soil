package key

import (
	"github.com/funtimecoding/soil/pkg/technitium/constant"
	"strings"
)

func Encode(v []*Key) string {
	if len(v) == 0 {
		return constant.ClearValue
	}

	var fields []string

	for _, k := range v {
		fields = append(fields, k.Name, k.Secret, k.Algorithm)
	}

	return strings.Join(fields, constant.RowSeparator)
}
