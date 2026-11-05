package option

import (
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/face"
)

func New() *Exporter {
	return &Exporter{
		Source: func() face.GitHubClient {
			return github.NewEnvironment()
		},
	}
}
