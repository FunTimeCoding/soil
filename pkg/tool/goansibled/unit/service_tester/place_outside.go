package service_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func (o *Tester) PlaceOutside() {
	system.WriteFile(
		filepath.Join(o.capture, constant.FixtureOutside),
		[]byte{},
		0o644,
	)
}
