package fix_tester

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func (o *Tester) Builds(directory string) {
	o.t.Helper()
	r := run.New().NoPanic()
	r.Directory = directory
	r.Execute(soil.Go, soil.Build, goModConstant.AllPackages)

	if r.Error != nil {
		o.t.Errorf("module does not build: %v\n%s", r.Error, r.ErrorString)
	}
}
