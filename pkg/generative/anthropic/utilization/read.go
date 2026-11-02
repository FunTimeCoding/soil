package utilization

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (c *Client) Read(
	x context.Context,
	token string,
) (*Result, error) {
	b, e := c.requester.Bytes(
		request.Get(constant.AnthropicUtilizationPath).
			WithContext(x).
			WithHeader(
				webConstant.Authorization,
				join.Space(webConstant.Bearer, token),
			),
	)

	if e != nil {
		return nil, e
	}

	result := Parse(b)

	if result == nil {
		return nil, unexpected.Format("utilization answer carries no limits")
	}

	return result, nil
}
