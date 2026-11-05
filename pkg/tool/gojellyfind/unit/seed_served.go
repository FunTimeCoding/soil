package unit

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/jellyfin/mock_client"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
)

func seedServed(c *mock_client.Client) {
	movie := item.Stub()
	movie.Identifier = "movie-1"
	movie.Name = "Alfa"
	movie.Type = "Movie"
	c.AddItem(movie)
	episode := item.Stub()
	episode.Identifier = "episode-1"
	episode.Name = "Bravo"
	episode.Type = "Episode"
	episode.SeriesName = "Charlie"
	c.AddItem(episode)
	track := item.Stub()
	track.Identifier = "track-1"
	track.Name = "Delta"
	track.Type = "Audio"
	c.AddItem(track)
	c.AddLibrary(library.New("library-1", "Echo", ""))
	s := session.Stub()
	s.Identifier = "session-1"
	s.DeviceName = "Foxtrot"
	s.PlayState = "Idle"
	s.SupportsRemote = true
	c.AddSession(s)
}
