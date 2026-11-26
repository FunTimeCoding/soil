package ladybird

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) endSession(session string) error {
	q := web.NewDelete(
		c.driverLocator(join.Slash([]string{constant.SessionRoute, session})),
	)
	q.Close = true
	r, e := c.pool.Client().Do(q)

	if e != nil {
		return fmt.Errorf("end session: %w", e)
	}

	errors.LogClose(r.Body)

	return nil
}
