package component

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
)

func Alert(message string) gomponents.Node {
	if message == "" {
		return gomponents.Group(nil)
	}

	return layout.Alert(message)
}
