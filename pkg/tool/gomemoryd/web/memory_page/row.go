package memory_page

import "github.com/funtimecoding/soil/pkg/web/component/opening_row"

func (p *Page) row() *opening_row.OpeningRow {
	m := p.view.Memory

	return opening(m.Identifier, cells(m.Name, m.Description, m.Tags))
}
