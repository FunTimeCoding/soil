package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func sightingToggle(unclaimed bool) gomponents.Node {
	if unclaimed {
		return html.A(
			html.Href(constant.SightingsPath),
			gomponents.Text(constant.EveryHost),
		)
	}

	query := url.Values{}
	query.Set(constant.UnclaimedParameter, constant.UnclaimedValue)
	locator := url.URL{
		Path:     constant.SightingsPath,
		RawQuery: query.Encode(),
	}

	return html.A(
		html.Href(locator.String()),
		gomponents.Text(constant.UnclaimedOnly),
	)
}
