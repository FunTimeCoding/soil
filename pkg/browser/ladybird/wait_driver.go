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
		constant.LadybirdDriverWait,
		constant.LadybirdDriverPoll,
		c.Reachable,
	)
	errors.PanicOnError(e)
}
