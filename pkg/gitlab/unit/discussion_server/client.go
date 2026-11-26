package discussion_server

import "github.com/funtimecoding/soil/pkg/gitlab"

func (s *Server) Client() *gitlab.Client {
	return s.client
}
