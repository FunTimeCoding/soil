package memory_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
)

func (p *Page) facts() gomponents.Node {
	m := p.view.Memory

	return fact.New().
		Row("Type", m.Type).
		Row("Scope", component.ScopeName(m.Scope)).
		RowNode("Created", stamp(m.CreatedAt)).
		RowNode("Updated", stamp(m.UpdatedAt)).
		Row("Active", fmt.Sprintf("%v", m.IsActive)).
		Node()
}
