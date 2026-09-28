package service

import (
	"github.com/ory/fosite"
	"net/http"
)

func (s *Service) NewAccessResponse(
	r *http.Request,
	accessRequest fosite.AccessRequester,
) (fosite.AccessResponder, error) {
	return s.provider.NewAccessResponse(r.Context(), accessRequest)
}
