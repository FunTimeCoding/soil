package export_test_helpers

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func exportedName(name string) string {
	if name == "" {
		return name
	}

	return join.Empty(strings.ToUpper(name[:1]), name[1:])
}
