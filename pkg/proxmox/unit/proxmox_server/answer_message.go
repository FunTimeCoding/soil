package proxmox_server

import "github.com/funtimecoding/soil/pkg/proxmox/unit/proxmox_server/answer"

func (s *Server) AnswerMessage(
	path string,
	line string,
	body string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.answers[path] = answer.New(line, body)
}
