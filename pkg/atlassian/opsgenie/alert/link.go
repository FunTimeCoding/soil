package alert

import "github.com/funtimecoding/soil/pkg/web/locator"

func link(
	identifier string,
	webHost string,
) string {
	return locator.New(
		webHost,
	).Path("/alert/detail/%s/details", identifier).String()
}
