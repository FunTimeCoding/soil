package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"maragu.dev/gomponents"
)

func tagText(tags []string) gomponents.Node {
	return gomponents.Text(join.Space(tags...))
}
