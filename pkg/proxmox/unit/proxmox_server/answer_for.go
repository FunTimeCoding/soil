package proxmox_server

import "github.com/funtimecoding/soil/pkg/strings/join"

func (s *Server) AnswerFor(
	method string,
	path string,
	line string,
	body string,
) {
	s.AnswerMessage(join.Space(method, path), line, body)
}
