package alert

import (
	"github.com/funtimecoding/soil/pkg/system"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
)

func (a *Alert) ResolveHost() string {
	if a.HostResolved() {
		return a.FullHost()
	}

	result := system.Lookup(a.FullHost())

	if len(result) == 0 {
		return systemConstant.UnknownHost
	}

	return result[0]
}
