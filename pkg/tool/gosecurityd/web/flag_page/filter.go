package flag_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func Filter(show string) gomponents.Node {
	chips := make([]gomponents.Node, 0, 3)

	for _, v := range [][2]string{
		{"Open", ""},
		{"Mitigated", "mitigated"},
		{"All", "all"},
	} {
		link := constant.FlagsPath

		if v[1] != "" {
			link = join.Empty(
				constant.FlagsPath,
				"?",
				url.Values{constant.ShowField: {v[1]}}.Encode(),
			)
		}

		c := chip.New(v[0]).Link(link)

		if v[1] == show {
			c.Current()
		}

		chips = append(chips, c.Node())
	}

	return html.Nav(chips...)
}
