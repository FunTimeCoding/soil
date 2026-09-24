package registry

import "github.com/funtimecoding/soil/pkg/measure/constant"

func NewDefault() *Registry {
	return New(constant.Languages...)
}
