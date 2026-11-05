package basic

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/bearer"
	"net/http"
)

func New(
	base *locator.Locator,
	token string,
) *Client {
	return &Client{
		requester: requester.New(base).
			WithAuthorizer(bearer.New(token)).
			WithHeader(webConstant.Accept, webConstant.Object).
			WithSubject("linkace").
			WithStatus(http.StatusBadRequest, errorsConstant.ValidationClass).
			WithStatus(
				http.StatusUnprocessableEntity,
				errorsConstant.ValidationClass,
			),
	}
}
