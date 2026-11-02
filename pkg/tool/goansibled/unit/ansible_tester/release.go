package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"path/filepath"
)

func (o *Tester) Release() {
	system.Remove(filepath.Join(o.capture, constant.FixtureHold))
}
