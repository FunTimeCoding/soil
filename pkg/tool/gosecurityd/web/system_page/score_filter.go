package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/chip"
	"maps"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func ScoreFilter(
	path string,
	kept url.Values,
	current float64,
	options [][2]string,
) gomponents.Node {
	chips := make([]gomponents.Node, 0, len(options))

	for _, v := range options {
		q := url.Values{}
		maps.Copy(q, kept)
		value := v[1]

		if value == "" {
			value = ScoreValue(0)
		} else {
			q.Set(constant.ScoreField, value)
		}

		c := chip.New(v[0]).Link(Query(path, q))

		if value == ScoreValue(current) {
			c.Current()
		}

		chips = append(chips, c.Node())
	}

	return html.Nav(chips...)
}
