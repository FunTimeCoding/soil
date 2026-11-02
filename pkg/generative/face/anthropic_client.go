package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization"
)

type AnthropicClient interface {
	Close()
	Reachable(x context.Context) error
	Read(x context.Context, token string) (*utilization.Result, error)
}
