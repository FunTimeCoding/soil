package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func (o *Tester) putKernel(
	name string,
	release string,
	kernel string,
	packages ...installed.Installed,
) {
	o.t.Helper()
	v := system.New(name, release, 0)
	v.Kernel = kernel
	assert.FatalOnError(o.t, o.System().Put(v, packages))
}
