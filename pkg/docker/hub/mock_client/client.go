package mock_client

import "github.com/funtimecoding/soil/pkg/docker/hub/tag"

type Client struct {
	Failure error
	Closed  int
	tags    map[string][]*tag.Tag
}
