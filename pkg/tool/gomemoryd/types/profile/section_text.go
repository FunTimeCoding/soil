package profile

import "github.com/funtimecoding/soil/pkg/strings/join"

func sectionText(
	heading string,
	body []string,
) string {
	return join.Empty(Opening(heading), join.NewLine(body))
}
