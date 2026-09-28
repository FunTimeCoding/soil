package service

import (
	"github.com/ory/fosite"
	"net/http"
)

func (s *Service) WriteAuthorizeResponse(
	r *http.Request,
	w http.ResponseWriter,
	authorizeRequest fosite.AuthorizeRequester,
	authorizeResponse fosite.AuthorizeResponder,
) {
	s.provider.WriteAuthorizeResponse(
		r.Context(),
		w,
		authorizeRequest,
		authorizeResponse,
	)
}
