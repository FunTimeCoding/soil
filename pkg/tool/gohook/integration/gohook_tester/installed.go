package gohook_tester

import (
	"github.com/funtimecoding/soil/pkg/git/repository_tester"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"github.com/funtimecoding/soil/pkg/tool/gohook/configuration"
	"github.com/funtimecoding/soil/pkg/tool/gohook/constant"
	"testing"
)

func Installed(
	t *testing.T,
	hooks string,
) string {
	t.Helper()
	clone := repository_tester.New(t).Clone
	repository_tester.NewAt(clone).Write(constant.RootFile, hooks)
	gohook.Install(configuration.Load(clone))

	return clone
}
