package convert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/server"
)

func ConfluencePages(v []*search_result.Result) []*server.ConfluencePage {
	result := make([]*server.ConfluencePage, 0, len(v))

	for _, r := range v {
		p := &server.ConfluencePage{
			Identifier: r.Raw.Identifier,
			Title:      r.Name,
		}

		if r.Raw.Type != "" {
			p.Type = &r.Raw.Type
		}

		if r.Raw.Status != "" {
			p.Status = &r.Raw.Status
		}

		if r.Raw.Links.Self != "" {
			p.Link = &r.Raw.Links.Self
		}

		result = append(result, p)
	}

	return result
}
