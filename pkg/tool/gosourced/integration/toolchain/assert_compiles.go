package toolchain

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
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
	r.Start(soil.Go, "vet", goModConstant.AllPackages)

	if r.Error != nil {
		t.Fatalf("vet: %s", r.ErrorString)
	}
}
