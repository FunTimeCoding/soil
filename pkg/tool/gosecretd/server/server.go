package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	secret "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"
)

type Server struct {
	service  secret.Secrets
	reporter face.Reporter
}
