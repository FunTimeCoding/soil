package memory_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/memory_list"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func scopes(l *memory_list.List) gomponents.Node {
	if len(l.Scopes) < 2 {
		return gomponents.Group(nil)
	}

	all := l.Scope == "" || l.Scope == constant.AllScope
	chips := []gomponents.Node{scopeChip(l, constant.AllScope, "", all)}

	for _, scope := range l.Scopes {
		label := scope

		if scope == "" {
			label = constant.DefaultScope
		}

		chips = append(chips, scopeChip(l, label, label, l.Scope == label))
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Text("Scope:"),
		gomponents.Group(chips),
	)
}
