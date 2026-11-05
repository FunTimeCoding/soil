package option

import "github.com/funtimecoding/soil/pkg/openbao/face"

type Secret struct {
	Address       string
	MetricAddress string
	ServiceTokens []string
	Source        func() face.OpenbaoClient
}
