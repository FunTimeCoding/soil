package service

import "net/http"

func (s *Service) Bridge() http.Handler {
	return s.client
}
