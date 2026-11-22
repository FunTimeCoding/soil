package service_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func (o *Tester) DefineHostCommand() {
	system.WriteFile(
		filepath.Join(
			o.Option.CloneDirectory,
			o.Option.AnsibleDirectory,
			constant.HostCommandFile,
		),
		[]byte(constant.FixtureHostCommand),
		0o644,
	)
}
