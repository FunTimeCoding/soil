package page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
)

func PrintBody(b response.Body) {
	if false {
		fmt.Printf("    Storage: %s\n", b.Storage.Value)
	}

	if false {
		fmt.Printf("    Text: %s\n", ToText(b.Storage.Value))
	}

	fmt.Printf("    Markdown: %s\n", bodyToMarkdown(b))
}
