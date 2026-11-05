package option

import "github.com/funtimecoding/soil/pkg/opnsense/face"

type Opnsense struct {
	Address       string
	ServiceTokens []string
	Source        func() face.OpnsenseClient
}
