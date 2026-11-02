package goband

import (
	"github.com/funtimecoding/soil/pkg/band"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	gobanddConstant "github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func pendingRedirection(v *client.Status) []string {
	var result []string

	if !v.Listener {
		result = append(result, gobanddConstant.ListenerWrite)
	}

	if v.ScreenService != band.ScreenServiceStateName(
		bandConstant.ScreenServiceEnabled,
	) &&
		v.ScreenService != band.ScreenServiceStateName(
			bandConstant.ScreenServiceOffline,
		) {
		result = append(result, gobanddConstant.ScreenServiceWrite)
	}

	return result
}
