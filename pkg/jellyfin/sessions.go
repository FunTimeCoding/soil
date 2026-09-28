package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/response"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
)

func (c *Client) Sessions() ([]*session.Session, error) {
	var out []response.Session
	e := c.get("/Sessions", nil, &out)

	if e != nil {
		return nil, e
	}

	result := make([]*session.Session, 0, len(out))

	for i := range out {
		r := &out[i]
		playState := "stopped"
		position := ""

		if r.PlayState != nil && r.NowPlayingItem != nil {
			if r.PlayState.IsPaused {
				playState = "paused"
			} else {
				playState = "playing"
			}

			position = formatTicks(r.PlayState.PositionTicks)
		}

		nowPlayingName := ""
		nowPlayingType := ""
		nowPlayingIdentifier := ""

		if r.NowPlayingItem != nil {
			nowPlayingName = r.NowPlayingItem.Name
			nowPlayingType = r.NowPlayingItem.Type
			nowPlayingIdentifier = r.NowPlayingItem.Identifier
		}

		result = append(
			result,
			session.New(
				r.Identifier,
				r.DeviceName,
				r.DeviceIdentifier,
				r.Client,
				r.UserName,
				nowPlayingName,
				nowPlayingType,
				nowPlayingIdentifier,
				playState,
				position,
				r.SupportsRemoteControl,
			),
		)
	}

	return result, nil
}
