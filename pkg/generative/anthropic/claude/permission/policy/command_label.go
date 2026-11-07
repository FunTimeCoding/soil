package policy

import "github.com/funtimecoding/soil/pkg/strings/join"

func (w *Walk) commandLabel(words []string) string {
	if len(words) < 2 {
		return words[0]
	}

	for _, r := range w.rules {
		if len(r.Words) > 1 && r.Words[0] == words[0] {
			return join.Space(words[0], words[1])
		}
	}

	return words[0]
}
