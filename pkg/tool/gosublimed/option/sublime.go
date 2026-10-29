package option

import "github.com/funtimecoding/soil/pkg/tool/gosublimed/face"

type Sublime struct {
	Address       string
	ServiceTokens []string
	Source        func() face.SublimeSource
}
