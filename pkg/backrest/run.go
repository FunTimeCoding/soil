package backrest

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (c *Client) run(
	x context.Context,
	arguments ...string,
) (string, error) {
	r := run.New().NoPanic().WithContext(x)
	output := r.Start(append(c.command, arguments...)...)

	return output, r.Error
}
