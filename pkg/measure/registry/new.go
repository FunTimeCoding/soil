package registry

import "github.com/funtimecoding/soil/pkg/measure/language"

func New(v ...*language.Language) *Registry {
	return &Registry{languages: v}
}
