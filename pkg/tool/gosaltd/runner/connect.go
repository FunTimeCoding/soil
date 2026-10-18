package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
)

func (r *Runner) connect(
	x context.Context,
	c *salt.Client,
) bool {
	e := c.Login(x)

	if x.Err() != nil {
		return false
	}

	r.reacher.Handle(
		constant.SaltPeer,
		e,
		func(y *edge.Edge) { r.logger.Plain("salt API %s", y) },
		func(f error) { r.reporter.CaptureException(f) },
	)

	return e == nil
}
