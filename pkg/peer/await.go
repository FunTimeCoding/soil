package peer

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"github.com/funtimecoding/soil/pkg/peer/retry"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Await(
	x context.Context,
	l *logger.Logger,
	name string,
	address string,
	probe func(context.Context) error,
) error {
	if _, e := retry.Until(x, constant.Wait, constant.Poll, probe); e != nil {
		return fmt.Errorf("%s not reachable at %s: %w", name, address, e)
	}

	l.Structured(join.Empty(name, "_connected"), "address", address)

	return nil
}
