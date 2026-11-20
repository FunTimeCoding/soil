package rewrite_tester

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (o *Tester) Compiles(directory string) {
	o.t.Helper()
	r := run.New().NoPanic()
	r.Directory = directory
	r.Start(soil.Go, "vet", goModConstant.AllPackages)

	if r.Error != nil {
		o.t.Fatalf("vet: %s", r.ErrorString)
	}
}
