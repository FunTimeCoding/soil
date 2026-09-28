package service

import (
	"github.com/ory/fosite"
	"net/http"
)

func (s *Service) WriteAuthorizeError(
	r *http.Request,
	w http.ResponseWriter,
	authorizeRequest fosite.AuthorizeRequester,
	e error,
) {
	s.provider.WriteAuthorizeError(r.Context(), w, authorizeRequest, e)
}
