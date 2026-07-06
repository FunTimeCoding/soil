package alert

import "github.com/funtimecoding/soil/pkg/text/markdown"

func (a *Alert) MarkdownLink() string {
	return markdown.Link(a.Name, a.Link)
}
