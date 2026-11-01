package toolchain

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"testing"
)

func assertCompiles(
	t *testing.T,
	directory string,
) {
	t.Helper()
	r := run.New().NoPanic()
	r.Directory = directory
	r.Start(constant.Go, "vet", "./...")

	if r.Error != nil {
		t.Fatalf("vet: %s", r.ErrorString)
	}
}
