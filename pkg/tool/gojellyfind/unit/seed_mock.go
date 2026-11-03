package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func seedMock(c *mock_client.Client) {
	movie := &client.Item{Id: "movie-1", Name: "Alfa", Type: "Movie"}
	episode := &client.Item{
		Id:         "episode-1",
		Name:       "Bravo",
		Type:       "Episode",
		SeriesName: new("Charlie"),
	}
	track := &client.Item{Id: "track-1", Name: "Delta", Type: "Audio"}
	c.ItemList = []*client.Item{movie, episode, track}
	c.Episodes["series-1"] = []*client.Item{episode}
	c.Tracks["album-1"] = []*client.Item{track}
	c.LibraryList = []*client.Library{{Id: "library-1", Name: "Echo"}}
	c.SessionList = []*client.Session{
		{
			Id:             "session-1",
			DeviceName:     "Foxtrot",
			PlayState:      "Idle",
			SupportsRemote: true,
		},
	}
}
