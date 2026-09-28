package service

import (
	"github.com/ory/fosite"
	"net/http"
)

func (s *Service) WriteAccessResponse(
	r *http.Request,
	w http.ResponseWriter,
	accessRequest fosite.AccessRequester,
	accessResponse fosite.AccessResponder,
) {
	s.provider.WriteAccessResponse(
		r.Context(),
		w,
		accessRequest,
		accessResponse,
	)
}
