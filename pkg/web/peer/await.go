package peer

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/retry"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func Await(
	l *logger.Logger,
	name string,
	address string,
	probe func() error,
) error {
	if _, e := retry.Until(constant.PeerWait, constant.PeerPoll, probe); e != nil {
		return fmt.Errorf("%s not reachable at %s: %w", name, address, e)
	}

	l.Structured(join.Empty(name, "_connected"), "address", address)

	return nil
}
