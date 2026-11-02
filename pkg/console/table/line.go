package table

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func (t *Table) line(
	b *strings.Builder,
	values []string,
) {
	for i, c := range t.columns {
		if i > 0 {
			b.WriteString(stringsConstant.DoubleSpace)
		}

		var v string

		if i < len(values) {
			v = values[i]
		}

		b.WriteString(pad(v, c.Width, c.Right))
	}

	b.WriteString(stringsConstant.Unix)
}
