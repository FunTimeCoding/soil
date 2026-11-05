package mock_client

import "github.com/funtimecoding/soil/pkg/github/run"

type Client struct {
	Failure error
	Closed  int
	runs    map[string][]*run.Run
}
