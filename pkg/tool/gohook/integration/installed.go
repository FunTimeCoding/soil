package integration

import (
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"testing"
)

func installed(
	t *testing.T,
	hooks string,
) string {
	t.Helper()
	clone := repository(t)
	write(clone, constant.RootFile, hooks)
	gohook.Install(configuration.Load(clone))

	return clone
}
