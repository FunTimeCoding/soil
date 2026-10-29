package reachable_probe

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"io"
	"net/http"
)

func Reach(
	x context.Context,
	c *http.Client,
	base *locator.Locator,
	token string,
) error {
	r, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		base.Copy().Path(constant.ReachablePath).String(),
		nil,
	)

	if e != nil {
		return fmt.Errorf("reachable request: %w", e)
	}

	web.Bearer(r, token)
	response, f := c.Do(r)

	if f != nil {
		return f
	}

	defer errors.LogClose(response.Body)

	if response.StatusCode == http.StatusNoContent {
		return nil
	}

	body, g := io.ReadAll(response.Body)

	if g != nil {
		return fmt.Errorf("read reachable: %w", g)
	}

	return refusal.New("reachable", response, body)
}
