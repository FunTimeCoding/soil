package registry

import "github.com/funtimecoding/soil/pkg/measure/language"

func (r *Registry) Languages() []*language.Language {
	return r.languages
}
