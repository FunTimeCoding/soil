package basic

import (
	"github.com/funtimecoding/soil/pkg/atlassian/classify_refusal"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
)

func newRequester(
	base *locator.Locator,
	user string,
	token string,
) *requester.Requester {
	return requester.New(base).
		WithAuthorizer(basic.New(user, token)).
		WithHeader(webConstant.Accept, webConstant.Object).
		WithRefusal(classify_refusal.Refusal)
}
