package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/usage_result"
)

type AnthropicClient interface {
	Close()
	Reachable(x context.Context) error
	Read(x context.Context, token string) (*usage_result.Result, error)
}
