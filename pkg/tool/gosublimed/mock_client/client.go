package mock_client

import "github.com/funtimecoding/soil/pkg/sublime/view"

type Client struct {
	views  []*view.View
	nextID int
}
