package service

import (
	"github.com/ory/fosite"
	"net/http"
)

func (s *Service) NewAuthorizeResponse(
	r *http.Request,
	q fosite.AuthorizeRequester,
	i fosite.Session,
) (fosite.AuthorizeResponder, error) {
	q.SetSession(i)

	return s.provider.NewAuthorizeResponse(r.Context(), q, i)
}
