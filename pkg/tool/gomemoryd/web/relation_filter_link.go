package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func relationFilterLink(
	label string,
	value string,
	current string,
) gomponents.Node {
	if value == current {
		return html.Strong(gomponents.Text(label))
	}

	target := constant.RelationsPath

	if value != "" {
		params := url.Values{}
		params.Set(constant.Type, value)
		target = fmt.Sprintf(
			"%s?%s",
			constant.RelationsPath,
			params.Encode(),
		)
	}

	return html.A(
		gomponents.Attr("href", target),
		gomponents.Text(label),
	)
}
