package service

import "net/http"

func (s *Service) Subject(r *http.Request) string {
	return s.authorization.Subject(r)
}
