package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"sync"
)

type Client struct {
	mutex       sync.Mutex
	ItemList    []*client.Item
	Episodes    map[string][]*client.Item
	Tracks      map[string][]*client.Item
	LibraryList []*client.Library
	SessionList []*client.Session
	Played      map[string][]client.PlayJSONRequestBody
	Commanded   map[string][]client.PlaybackCommandJSONRequestBody
	Volumes     map[string][]int
	Failure     error
	Closed      int
	failures    *failure.Failure
}
