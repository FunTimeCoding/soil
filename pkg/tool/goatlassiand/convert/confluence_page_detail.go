package convert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func ConfluencePageDetail(p *page.Page) *server.ConfluencePageDetail {
	result := &server.ConfluencePageDetail{
		Identifier: p.Identifier,
		Title:      p.Name,
	}

	if p.Link != "" {
		result.Link = &p.Link
	}

	if p.Raw.SpaceIdentifier != "" {
		result.SpaceIdentifier = &p.Raw.SpaceIdentifier
	}

	if p.Raw.Body.Storage.Value != "" {
		result.Body = new(page.ToMarkdown(p.Raw.Body.Storage.Value))
	}

	return result
}
