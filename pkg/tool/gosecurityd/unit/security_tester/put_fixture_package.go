package security_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
)

func (o *Tester) PutFixturePackage(
	name string,
	release string,
	sourceVersion string,
	candidate string,
) {
	o.t.Helper()
	assert.FatalOnError(
		o.t,
		o.System().Put(
			system.New(name, release, 0),
			[]installed.Installed{
				*installed.New(
					"libtango5",
					sourceVersion,
					"tango-utils",
					sourceVersion,
					candidate,
				),
			},
		),
	)
}
