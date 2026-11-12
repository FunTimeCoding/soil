package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/usage_result"

type Client struct {
	Failure error
	Result  *usage_result.Result
	Reads   int
	Closed  int
}
