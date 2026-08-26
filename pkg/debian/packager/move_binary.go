package packager

import (
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/join"
	"path/filepath"
)

func (p *Packager) MoveBinary() {
	system.MakeDirectory(p.BinaryRoot)
	source := p.ExecutablePath

	if !runtime.RunningFromBinary() {
		source = filepath.Join(system.WorkDirectory(), p.ExecutablePath)
	}

	system.Move(source, join.Absolute(p.BinaryRoot, p.ExecutableName))
}
