package option

import "github.com/funtimecoding/soil/pkg/linkace/face"

type LinkAce struct {
	Address       string
	ServiceTokens []string
	Source        func() face.Client
}
