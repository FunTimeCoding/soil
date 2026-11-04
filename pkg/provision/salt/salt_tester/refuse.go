package salt_tester

import "fmt"

func (s *Server) Refuse(
	status int,
	paragraph string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.gateway = status
	s.page = fmt.Sprintf("<html><p>%s</p></html>", paragraph)
}
