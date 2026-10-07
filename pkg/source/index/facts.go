package index

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
)

func Facts[T any](
	w *Workspace,
	k *kind.Kind,
) map[string]T {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.ensure()
	values, registered := w.facts[k.Name]

	if !registered {
		panic(
			fmt.Sprintf("kind %s is not registered with this workspace", k.Name),
		)
	}

	result := make(map[string]T, len(values))

	for path, v := range values {
		result[path] = v.(T)
	}

	return result
}
