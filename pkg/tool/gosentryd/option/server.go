package option

import "github.com/funtimecoding/soil/pkg/errors/sentry/face"

type Server struct {
	Address       string
	ServiceTokens []string
	Organization  string
	Host          string
	Source        func() face.Client
}
