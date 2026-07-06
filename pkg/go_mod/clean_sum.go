package go_mod

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func cleanSum(
	mod string,
	version string,
) {
	path := "go.sum"

	if !system.FileExists(path) {
		return
	}

	content := system.ReadFile(system.WorkDirectory(), path)
	prefix := fmt.Sprintf("%s %s", mod, version)
	var lines []string

	for _, line := range strings.Split(content, "\n") {
		if !strings.HasPrefix(line, prefix) {
			lines = append(lines, line)
		}
	}

	system.SaveFile(path, strings.Join(lines, "\n"))
}
