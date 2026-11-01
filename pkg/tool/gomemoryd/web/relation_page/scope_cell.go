package relation_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web/component"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func scopeCell(scope string) gomponents.Node {
	return html.Td(html.Small(gomponents.Text(component.ScopeName(scope))))
}
