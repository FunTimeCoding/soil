package page

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"

func New(
	p *response.Page,
	host string,
) *Page {
	return &Page{
		Identifier:     p.Identifier,
		Name:           p.Title,
		Status:         p.Status,
		Body:           p.Body.Storage.Value,
		VersionMessage: p.Version.Message,
		Link:           link(p, host, false),
		TinyLink:       link(p, host, true),
		Raw:            p,
	}
}
