package loader

import "github.com/funtimecoding/soil/pkg/notation"

func (l *Loader) ToMap() map[string]map[string]string {
	var result = map[string]map[string]string{}

	for k, v := range l.contents {
		var e = map[string]string{}
		notation.MustDecode(v, &e, false)
		result[k] = e
	}

	return result
}
