package fact

import (
	"github.com/funtimecoding/soil/pkg/web/component/fact/row"
	"maragu.dev/gomponents"
)

func (f *Fact) RowNode(
	label string,
	value gomponents.Node,
) *Fact {
	f.rows = append(f.rows, row.New(label, value))

	return f
}
