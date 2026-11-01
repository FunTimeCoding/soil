package history_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/history_filter"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"slices"
)

func Filters(active []string) gomponents.Node {
	chips := []gomponents.Node{}

	for _, f := range []*history_filter.Filter{
		history_filter.New("Completions", constant.Complete),
		history_filter.New("Summaries", constant.Summarize),
		history_filter.New("Moments", constant.Moment),
		history_filter.New("Updates", constant.Update),
	} {
		c := chip.New(f.Label)

		if slices.Contains(active, f.Kind) {
			c.Link(
				Path(
					slices.DeleteFunc(
						slices.Clone(active),
						func(k string) bool { return k == f.Kind },
					),
					1,
				),
			).Current()
		} else {
			c.Link(Path(append(slices.Clone(active), f.Kind), 1))
		}

		chips = append(chips, c.Node())
	}

	return html.Div(
		html.Class(webConstant.ClusterClass),
		gomponents.Group(chips),
	)
}
