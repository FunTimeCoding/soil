package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
)

type Client interface {
	GetEntries(
		x context.Context,
		p *client.GetEntriesParams,
	) ([]client.EntryResponse, error)
	PostEntry(
		x context.Context,
		b client.PostEntryRequest,
	) (*client.EntryResponse, error)
	UpdateEntry(
		x context.Context,
		identifier int,
		b client.PostEntryRequest,
	) (*client.EntryResponse, error)
	DeleteEntry(
		x context.Context,
		identifier int,
	) error
	GetStatus(x context.Context) (*client.StatusResponse, error)
	Reachable(x context.Context) error
	Address() string
	Close()
}
