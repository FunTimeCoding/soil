package option

import "github.com/funtimecoding/soil/pkg/habitica/face"

type Habitica struct {
	Address       string
	ServiceTokens []string
	Host          string
	Source        func() face.Client
}
