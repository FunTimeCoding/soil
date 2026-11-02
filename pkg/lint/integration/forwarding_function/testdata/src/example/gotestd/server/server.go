package server

import "example/gotestd/service/loan"

type Server struct {
	loan *loan.Service
}

func (s *Server) Renew(identifier int) error {
	return s.loan.Renew(identifier)
}
