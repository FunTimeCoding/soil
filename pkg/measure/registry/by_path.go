package registry

import (
	"github.com/funtimecoding/soil/pkg/measure/language"
	"path/filepath"
	"slices"
	"strings"
)

func (r *Registry) ByPath(path string) *language.Language {
	name := filepath.Base(path)

	for _, l := range r.languages {
		for _, s := range l.Suffixes {
			if strings.HasSuffix(name, s) {
				return l
			}
		}
	}

	for _, l := range r.languages {
		if slices.Contains(l.Filenames, name) {
			return l
		}
	}

	extension := filepath.Ext(name)

	for _, l := range r.languages {
		if extension != "" && slices.Contains(l.Extensions, extension) {
			return l
		}
	}

	return nil
}
