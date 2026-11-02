package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

type Client interface {
	GetEpisodes(
		x context.Context,
		identifier string,
	) ([]*client.Item, error)
	GetItem(
		x context.Context,
		identifier string,
	) (*client.Item, error)
	GetTracks(
		x context.Context,
		identifier string,
	) ([]*client.Item, error)
	ListLibraries(x context.Context) ([]*client.Library, error)
	ListSessions(x context.Context) ([]*client.Session, error)
	Play(
		x context.Context,
		identifier string,
		b client.PlayJSONRequestBody,
	) error
	PlaybackCommand(
		x context.Context,
		identifier string,
		b client.PlaybackCommandJSONRequestBody,
	) error
	SearchItems(
		x context.Context,
		p *client.SearchItemsParams,
	) (*client.SearchResult, error)
	SetVolume(
		x context.Context,
		identifier string,
		b client.SetVolumeJSONRequestBody,
	) error
	Reachable(x context.Context) error
	Close()
}
