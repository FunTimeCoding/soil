package hub

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

func NewWithBase(base *locator.Locator) *Client {
	return &Client{
		requester: requester.New(base).WithHeader(
			webConstant.Accept,
			webConstant.Object,
		).WithRefusal(classifyRefusal),
	}
}
