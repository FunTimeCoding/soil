package base

import "github.com/funtimecoding/soil/pkg/directory/face"

func WithDirectory(d face.DirectoryClient) Option {
	return func(o *Server) {
		o.directory = d
	}
}
