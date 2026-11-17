package base

import "github.com/funtimecoding/soil/pkg/jellyfin/face"

func WithSource(c face.JellyfinClient) Option {
	return func(s *Server) {
		s.source = c
	}
}
