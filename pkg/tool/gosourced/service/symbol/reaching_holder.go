package symbol

import "github.com/funtimecoding/soil/pkg/source/index"

func reachingHolder(
	w *index.Workspace,
	holders []string,
	target string,
) (string, bool) {
	for _, h := range holders {
		if h == target {
			return h, true
		}

		dependencies, complete := w.Dependencies(h)

		if dependencies[target] || !complete {
			return h, true
		}
	}

	return "", false
}
