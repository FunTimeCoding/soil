package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func (p *Page) Editor(alert string) gomponents.Node {
	m := p.view.Memory
	r := p.row()

	return r.Detail(
		len(cells(m.Name, m.Description, m.Tags)),
		entry.New(actionPath(m.Identifier, constant.EditAction), "Save").
			Swap(r.Target(), webConstant.SwapOuter).
			Alert(alert).
			Text(constant.MemoryName, "Name", m.Name).
			Required().
			Text(constant.Description, "Description", m.Description).
			TextArea(constant.Content, "Content", m.Content).
			Lines("16").
			Beside(
				load.New(
					actionPath(m.Identifier, constant.OpenAction),
					"Cancel",
				).
					Swap(r.Target(), webConstant.SwapOuter).
					Button().
					Class("outline secondary").
					Node(),
			).
			Node(),
	).Response()
}
