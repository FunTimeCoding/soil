package face

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/types/search"
)

type Jellyfin interface {
	Libraries() ([]*library.Library, error)
	SearchItems(
		term string,
		types []string,
		page int,
		perPage int,
	) (*search.Result, error)
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
}
