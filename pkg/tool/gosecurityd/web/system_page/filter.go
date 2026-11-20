package system_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func Filter(current string) gomponents.Node {
	chips := make([]gomponents.Node, 0, 4)

	for _, v := range [][2]string{
		{"All", ""},
		{"Supported", "supported"},
		{"Tracked", "tracked"},
		{"Untracked", "untracked"},
	} {
		link := constant.SystemsPath

		if v[1] != "" {
			link = join.Empty(
				constant.SystemsPath,
				"?",
				url.Values{constant.ShowField: {v[1]}}.Encode(),
			)
		}

		c := chip.New(v[0]).Link(link)

		if v[1] == current {
			c.Current()
		}

		chips = append(chips, c.Node())
	}

	return html.Nav(chips...)
}
