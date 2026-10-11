package packager

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/linux/debian"
)

func (p *Packager) SaveUnit() {
	system.MakeDirectory(p.UnitRoot)
	system.SaveFile(p.unitPath(), debian.RenderUnit(p.ExecutableName))
}
