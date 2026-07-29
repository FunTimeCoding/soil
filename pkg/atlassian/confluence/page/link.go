package page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func link(
	p *response.Page,
	host string,
	tiny bool,
) string {
	if tiny {
		if p.Links.WebUI != "" {
			return locator.New(
				host,
			).Base(constant.ConfluenceWiki).Path(p.Links.TinyUI).String()
		}
	} else {
		if p.Links.WebUI != "" {
			return locator.New(
				host,
			).Base(constant.ConfluenceWiki).Path(p.Links.WebUI).String()
		}
	}

	return ""
}
