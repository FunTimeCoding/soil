package go_mod

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"strings"
)

func cleanSum(
	mod string,
	version string,
) {
	path := constant.SumFile

	if !system.FileExists(path) {
		return
	}

	content := system.ReadFile(system.WorkDirectory(), path)
	prefix := fmt.Sprintf("%s %s", mod, version)
	var lines []string

	for line := range strings.SplitSeq(content, stringConstant.Unix) {
		if !strings.HasPrefix(line, prefix) {
			lines = append(lines, line)
		}
	}

	system.SaveFile(path, join.NewLine(lines))
}
