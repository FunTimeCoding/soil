package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func (o *Tester) Hold() {
	system.WriteFile(
		filepath.Join(o.capture, constant.FixtureHold),
		[]byte{},
		0o644,
	)
	o.t.Cleanup(o.Release)
}
