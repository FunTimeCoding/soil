package server

import "net/http"

func (s *Server) token(
	w http.ResponseWriter,
	r *http.Request,
) {
	accessRequest, e := s.service.NewAccessRequest(r)

	if e != nil {
		s.service.WriteAccessError(r, w, e)

		return
	}

	accessResponse, e := s.service.NewAccessResponse(r, accessRequest)

	if e != nil {
		s.service.WriteAccessError(r, w, e)

		return
	}

	s.service.WriteAccessResponse(r, w, accessRequest, accessResponse)
}
