package toolchain

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func assertBuilds(
	t *testing.T,
	directory string,
) {
	t.Helper()
	r := run.New().NoPanic()
	r.Directory = directory
	r.Execute(constant.Go, constant.Build, "./...")

	if r.Error != nil {
		t.Errorf("module does not build: %v\n%s", r.Error, r.ErrorString)
	}
}
