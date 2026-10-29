package option

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry"
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
)

func New() *Server {
	return &Server{
		Source: func() face.Client {
			return sentry.NewEnvironment()
		},
	}
}
