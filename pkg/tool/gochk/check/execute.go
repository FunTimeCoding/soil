package check

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func Execute(c []string) string {
	return strings.TrimSpace(run.New().Start(c...))
}
