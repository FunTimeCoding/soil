package registry

import (
	"github.com/funtimecoding/soil/pkg/measure/language"
	"strings"
)

func (r *Registry) ByName(name string) *language.Language {
	for _, l := range r.languages {
		if strings.EqualFold(l.Name, name) {
			return l
		}
	}

	return nil
}
