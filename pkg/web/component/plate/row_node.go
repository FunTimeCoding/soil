package plate

import (
	"github.com/funtimecoding/soil/pkg/web/component/fact/row"
	"maragu.dev/gomponents"
)

func (p *Plate) RowNode(
	label string,
	value gomponents.Node,
) *Plate {
	p.rows = append(p.rows, row.New(label, value))

	return p
}
