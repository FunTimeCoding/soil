package service

import "net/http"

func (s *Service) WriteAccessError(
	r *http.Request,
	w http.ResponseWriter,
	e error,
) {
	s.provider.WriteAccessError(r.Context(), w, nil, e)
}
