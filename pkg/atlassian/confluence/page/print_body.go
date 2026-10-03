package page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/console"
)

func PrintBody(b response.Body) {
	console.Format("    Markdown: %s\n", bodyToMarkdown(b))
}
