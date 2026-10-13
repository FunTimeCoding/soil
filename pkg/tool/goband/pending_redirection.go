package goband

import (
	"github.com/funtimecoding/soil/pkg/band"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func pendingRedirection(v *client.Status) []string {
	var result []string

	if !v.Listener {
		result = append(result, constant.ListenerWrite)
	}

	if v.ScreenService != band.ScreenServiceStateName(
		bandConstant.ScreenServiceEnabled,
	) &&
		v.ScreenService != band.ScreenServiceStateName(
			bandConstant.ScreenServiceOffline,
		) {
		result = append(result, constant.ScreenServiceWrite)
	}

	return result
}
