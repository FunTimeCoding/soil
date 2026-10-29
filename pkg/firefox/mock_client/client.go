package mock_client

import (
	"github.com/funtimecoding/soil/pkg/failure"
	"github.com/funtimecoding/soil/pkg/firefox/download"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/firefox/tab_group"
	"sync"
)

type Client struct {
	mutex              sync.Mutex
	failures           *failure.Failure
	tabs               []*tab.Tab
	downloads          []*download.Download
	groups             map[int]*tab_group.Group
	nextIdentifier     int
	groupIdentifier    int
	downloadIdentifier int
	Closed             int
}
