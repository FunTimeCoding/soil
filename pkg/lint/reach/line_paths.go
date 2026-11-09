package reach

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path/filepath"
	"strings"
)

func linePaths(line string) []string {
	line = strings.TrimSpace(line)

	if m := constant.SourceRenamed.FindStringSubmatch(line); m != nil {
		return []string{m[1], filepath.Join(filepath.Dir(m[1]), m[3])}
	}

	if m := constant.SourceLocated.FindStringSubmatch(line); m != nil {
		return []string{m[1]}
	}

	if m := constant.SourceEmptied.FindStringSubmatch(line); m != nil {
		return []string{m[1]}
	}

	return nil
}
