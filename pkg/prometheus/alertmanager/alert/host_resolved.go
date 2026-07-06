package alert

import "github.com/funtimecoding/soil/pkg/web/address"

func (a *Alert) HostResolved() bool {
	return !address.IsInternet(a.FullHost())
}
