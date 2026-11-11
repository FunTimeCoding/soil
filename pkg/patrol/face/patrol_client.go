package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/mdstat/array"
	"github.com/funtimecoding/soil/pkg/smartctl/response"
)

type PatrolClient interface {
	Arrays(x context.Context) ([]*array.Array, error)
	Drives(x context.Context) ([]*response.Report, error)
	Reachable(x context.Context) error
}
