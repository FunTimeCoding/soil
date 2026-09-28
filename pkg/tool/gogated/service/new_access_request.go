package service

import (
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/openid"
	"net/http"
)

func (s *Service) NewAccessRequest(
	r *http.Request,
) (fosite.AccessRequester, error) {
	return s.provider.NewAccessRequest(r.Context(), r, &openid.DefaultSession{})
}
