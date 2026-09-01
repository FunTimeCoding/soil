package page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/console"
)

func PrintBody(b response.Body) {
	if false {
		console.Format("    Storage: %s\n", b.Storage.Value)
	}

	if false {
		console.Format("    Text: %s\n", ToText(b.Storage.Value))
	}

	console.Format("    Markdown: %s\n", bodyToMarkdown(b))
}
