package fact

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) Section(label string) *Fact {
	f.rows = append(
		f.rows,
		html.Tr(
			html.Class(constant.SectionClass),
			html.Th(html.ColSpan("2"), gomponents.Text(label)),
		),
	)

	return f
}
