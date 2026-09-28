package service

import (
	"github.com/ory/fosite"
	"net/http"
)

func (s *Service) NewAuthorizeRequest(
	r *http.Request,
) (fosite.AuthorizeRequester, error) {
	return s.provider.NewAuthorizeRequest(r.Context(), r)
}
