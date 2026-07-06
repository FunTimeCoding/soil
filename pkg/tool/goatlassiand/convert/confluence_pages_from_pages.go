package convert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func ConfluencePagesFromPages(v []*page.Page) []*server.ConfluencePage {
	result := make([]*server.ConfluencePage, 0, len(v))

	for _, p := range v {
		a := &server.ConfluencePage{Identifier: p.Identifier, Title: p.Name}

		if p.Link != "" {
			a.Link = &p.Link
		}

		result = append(result, a)
	}

	return result
}
