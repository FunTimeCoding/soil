package entry_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func detailFields(e *entry.Entry) gomponents.Node {
	result := fact.New()

	if e.Service != "" {
		result.Row("Service", e.Service)
	}

	if e.Description != "" {
		result.RowNode("Description", html.Pre(gomponents.Text(e.Description)))
	}

	return result.Node()
}
