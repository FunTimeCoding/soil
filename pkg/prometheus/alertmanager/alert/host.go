package alert

import (
	"github.com/funtimecoding/soil/pkg/web/address"
	"github.com/funtimecoding/soil/pkg/web/host"
)

func (a *Alert) Host() string {
	result := a.FullHost()

	if address.IsInternet(result) {
		return result
	}

	return host.StripDomain(result)
}
