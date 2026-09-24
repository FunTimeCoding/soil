package table

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func (t *Table) line(
	b *strings.Builder,
	values []string,
) {
	for i, c := range t.columns {
		if i > 0 {
			b.WriteString(constant.DoubleSpace)
		}

		var v string

		if i < len(values) {
			v = values[i]
		}

		b.WriteString(pad(v, c.width, c.right))
	}

	b.WriteString(constant.Unix)
}
