package fact

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (f *Fact) Node() gomponents.Node {
	return html.Table(html.Class(constant.FactsClass), html.TBody(f.rows...))
}
