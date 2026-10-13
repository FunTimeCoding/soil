package mock_client

import (
	"github.com/funtimecoding/soil/pkg/firefox/download"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/types/tab_group"
)

type Client struct {
	tabs               []*tab.Tab
	downloads          []*download.Download
	groups             map[int]*tab_group.Group
	nextIdentifier     int
	groupIdentifier    int
	downloadIdentifier int
}
