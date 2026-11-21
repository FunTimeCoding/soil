package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func Filter(
	current string,
	releases []string,
	minimum float64,
) gomponents.Node {
	options := [][2]string{{"All", ""}}

	for _, r := range releases {
		options = append(options, [2]string{r, r})
	}

	options = append(options, [2]string{"Kernel", constant.KernelFilter})
	chips := make([]gomponents.Node, 0, len(options))

	for _, v := range options {
		q := url.Values{}

		if v[1] != "" {
			q.Set(constant.ShowField, v[1])
		}

		if minimum > 0 {
			q.Set(constant.ScoreField, system_page.ScoreValue(minimum))
		}

		c := chip.New(v[0]).Link(system_page.Query(constant.UpgradesPath, q))

		if v[1] == current {
			c.Current()
		}

		chips = append(chips, c.Node())
	}

	kept := url.Values{}

	if current != "" {
		kept.Set(constant.ShowField, current)
	}

	return gomponents.Group(
		[]gomponents.Node{
			html.Nav(chips...),
			system_page.ScoreFilter(
				constant.UpgradesPath,
				kept,
				minimum,
				constant.CriticalChips,
			),
		},
	)
}
