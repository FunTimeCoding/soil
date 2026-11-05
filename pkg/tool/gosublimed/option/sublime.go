package option

import "github.com/funtimecoding/soil/pkg/sublime/face"

type Sublime struct {
	Address       string
	ServiceTokens []string
	Source        func() face.SublimeClient
}
