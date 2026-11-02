package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"

type Client struct {
	Failure error
	Result  *utilization.Result
	Reads   int
	Closed  int
}
