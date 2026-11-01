package memory_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func TagChips(tags []string) gomponents.Node {
	if len(tags) == 0 {
		return gomponents.Group(nil)
	}

	chips := make([]gomponents.Node, 0, len(tags))

	for _, t := range tags {
		chips = append(chips, chip.New(t).Link(TagPath(t)).Node())
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Group(chips),
	)
}
