package utilization

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/bearer"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func Read(token string) (*Result, error) {
	b, e := requester.New(locator.New(constant.AnthropicServiceHost)).
		WithAuthorizer(bearer.New(token)).
		WithHeader(
			constant.AnthropicBetaHeader,
			constant.AnthropicUtilizationBeta,
		).
		Bytes(request.Absolute(constant.AnthropicUtilizationLink))

	if e != nil {
		return nil, e
	}

	result := Parse(b)

	if result == nil {
		return nil, unexpected.Format("utilization answer carries no limits")
	}

	return result, nil
}
