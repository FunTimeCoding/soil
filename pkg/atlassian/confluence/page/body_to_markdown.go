package page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/text"
	"github.com/funtimecoding/soil/pkg/text/option"
)

func bodyToMarkdown(b response.Body) string {
	if b.Storage.Value == "" {
		return ""
	}

	result := ToMarkdown(b.Storage.Value)
	p := option.New()
	p.AllowedBlankLines = 0
	p.NewlineAtEnd = false
	result = text.OptimizeWhitespace(result, p)

	return result
}
