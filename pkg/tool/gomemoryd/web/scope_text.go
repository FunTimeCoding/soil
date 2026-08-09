package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func scopeText(scope string) gomponents.Node {
	if scope == "" {
		scope = constant.DefaultScope
	}

	return html.Small(gomponents.Text(scope))
}
