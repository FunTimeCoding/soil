package plate_page

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func PagesSection(
	title string,
	pages []*page.Page,
) gomponents.Node {
	if len(pages) == 0 {
		return gomponents.Group(nil)
	}

	return gomponents.Group(
		[]gomponents.Node{html.H3(gomponents.Text(title)), pageTable(pages)},
	)
}
