package hub

import (
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"net/http"
)

func NewWithBase(base *locator.Locator) *Client {
	return &Client{
		requester: requester.New(base).
			WithHeader(webConstant.Accept, webConstant.Object).
			WithSubject("docker hub").
			WithStatus(http.StatusForbidden, errorsConstant.ForbiddenClass),
	}
}
