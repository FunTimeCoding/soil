package identity

import "github.com/funtimecoding/soil/pkg/strings/join"

func (t *Tool) RenderInstructions(conditions map[string]bool) string {
	var parts []string
	parts = append(parts, t.instructions)

	for _, p := range t.paragraphs {
		if conditions[p.Key] {
			parts = append(parts, p.Text)
		}
	}

	return join.Space(parts...)
}
