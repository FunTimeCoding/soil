package option

import "github.com/funtimecoding/soil/pkg/tool/gosecretd/face"

type Secret struct {
	Address       string
	MetricAddress string
	ServiceTokens []string
	Source        func() face.SecretSource
}
