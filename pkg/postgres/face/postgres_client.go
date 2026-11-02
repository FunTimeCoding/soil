package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/postgres/types/result"
)

type PostgresClient interface {
	Close()
	Query(
		x context.Context,
		sql string,
		a ...any,
	) (*result.Result, error)
	Reachable(x context.Context) error
}
