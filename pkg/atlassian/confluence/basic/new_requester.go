package basic

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
	"net/http"
)

func newRequester(
	base *locator.Locator,
	user string,
	token string,
) *requester.Requester {
	return requester.New(base).
		WithAuthorizer(basic.New(user, token)).
		WithHeader(webConstant.Accept, webConstant.Object).
		WithStatus(http.StatusForbidden, errorsConstant.ForbiddenClass).
		WithStatus(
			http.StatusUnprocessableEntity,
			errorsConstant.ValidationClass,
		)
}
