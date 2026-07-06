package alert

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func (a *Alert) ResolveHost() string {
	if a.HostResolved() {
		return a.FullHost()
	}

	result := system.Lookup(a.FullHost())

	if len(result) == 0 {
		return constant.UnknownHost
	}

	return result[0]
}
