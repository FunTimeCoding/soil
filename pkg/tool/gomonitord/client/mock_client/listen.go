package mock_client

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/constant"
	"io"
	"strings"
)

func (c *Client) Listen(context.Context) (io.ReadCloser, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return io.NopCloser(
		strings.NewReader(
			fmt.Sprintf(constant.EventFormat, notation.Marshal(c.ClaimList)),
		),
	), nil
}
