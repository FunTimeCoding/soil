package face

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
)

type JellyfinSource interface {
	Libraries() ([]*library.Library, error)
	SearchItems(
		term string,
		types []string,
		page int,
		perPage int,
	) ([]*item.Item, int, error)
	Item(identifier string) (*item.Item, error)
	Episodes(seriesIdentifier string) ([]*item.Item, error)
	Tracks(albumIdentifier string) ([]*item.Item, error)
	Sessions() ([]*session.Session, error)
	Play(
		sessionIdentifier string,
		itemIdentifiers []string,
		playCommand string,
		startPositionTicks int64,
	) error
	PlaybackCommand(
		sessionIdentifier string,
		command string,
		seekPositionTicks int64,
	) error
	SetVolume(
		sessionIdentifier string,
		level int,
	) error
	MustLibraries() []*library.Library
	MustSearchItems(
		term string,
		types []string,
		page int,
		perPage int,
	) ([]*item.Item, int)
	MustItem(identifier string) *item.Item
	MustEpisodes(seriesIdentifier string) []*item.Item
	MustTracks(albumIdentifier string) []*item.Item
	MustSessions() []*session.Session
	MustPlay(
		sessionIdentifier string,
		itemIdentifiers []string,
		playCommand string,
		startPositionTicks int64,
	)
	MustPlaybackCommand(
		sessionIdentifier string,
		command string,
		seekPositionTicks int64,
	)
	MustSetVolume(
		sessionIdentifier string,
		level int,
	)
}
