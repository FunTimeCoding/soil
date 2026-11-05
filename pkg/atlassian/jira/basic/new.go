package basic

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
	"net/http"
)

func New(
	base *locator.Locator,
	user string,
	token string,
) *Client {
	return &Client{
		requester: requester.New(base).
			WithAuthorizer(basic.New(user, token)).
			WithHeader(webConstant.Accept, webConstant.Object).
			WithStatus(http.StatusForbidden, errorsConstant.ForbiddenClass).
			WithStatus(
				http.StatusUnprocessableEntity,
				errorsConstant.ValidationClass,
			),
		base: base,
	}
}
