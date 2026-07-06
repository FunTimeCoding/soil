package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page/page_file"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/text"
	"github.com/funtimecoding/soil/pkg/text/option"
)

func (c *Client) Export(
	p *page.Page,
	path string,
) {
	markdown := page.ToMarkdown(p.Raw.Body.Storage.Value)
	o := option.New()
	o.AllowedBlankLines = 0
	o.NewlineAtEnd = false
	markdown = text.OptimizeWhitespace(markdown, o)
	system.SaveFile(path, page_file.New(p.Name, markdown).Encode())
}
