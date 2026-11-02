package utilization

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

func New(base *locator.Locator) *Client {
	return &Client{
		requester: requester.New(base).WithHeader(
			constant.AnthropicBetaHeader,
			constant.AnthropicUtilizationBeta,
		).WithRefusal(classifyRefusal),
	}
}
