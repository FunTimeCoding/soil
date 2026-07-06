package session_store

import "github.com/funtimecoding/soil/pkg/web"

func (s *Store) Create(address string) string {
	s.Lock()
	result := web.GenerateSession()
	s.sessions[result] = address
	s.Unlock()

	return result
}
