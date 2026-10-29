package option

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"

type Opnsense struct {
	Address       string
	ServiceTokens []string
	Source        func() face.OpnsenseSource
}
