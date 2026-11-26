package ladybird

import (
	"context"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/peer/retry"
)

func (c *Client) waitDriver() {
	_, e := retry.Until(
		context.Background(),
		constant.DriverWait,
		constant.DriverPoll,
		c.Reachable,
	)
	errors.PanicOnError(e)
}
