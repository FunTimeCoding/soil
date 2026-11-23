package option

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/chat/teams"
)

func New() *Server {
	return &Server{
		Source: func() face.TeamsClient {
			return teams.NewEnvironment()
		},
	}
}
