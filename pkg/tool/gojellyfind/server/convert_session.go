package server

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/server"
)

func convertSession(v *session.Session) *server.Session {
	result := &server.Session{
		Id:             v.Identifier,
		DeviceName:     v.DeviceName,
		DeviceId:       v.DeviceIdentifier,
		Client:         v.Client,
		PlayState:      v.PlayState,
		SupportsRemote: v.SupportsRemote,
	}

	if v.UserName != "" {
		result.UserName = new(v.UserName)
	}

	if v.NowPlayingName != "" {
		result.NowPlayingName = new(v.NowPlayingName)
	}

	if v.NowPlayingType != "" {
		result.NowPlayingType = new(v.NowPlayingType)
	}

	if v.NowPlayingIdentifier != "" {
		result.NowPlayingId = new(v.NowPlayingIdentifier)
	}

	if v.Position != "" {
		result.Position = new(v.Position)
	}

	return result
}
