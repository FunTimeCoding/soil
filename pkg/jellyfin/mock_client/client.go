package mock_client

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
	"github.com/funtimecoding/soil/pkg/jellyfin/library"
	"github.com/funtimecoding/soil/pkg/jellyfin/session"
)

type Client struct {
	Failure   error
	Closed    int
	items     []*item.Item
	libraries []*library.Library
	sessions  []*session.Session
}
