package packager

import (
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/system"
)

func (p *Packager) SaveUnit() {
	system.MakeDirectory(p.UnitRoot)
	system.SaveFile(p.unitPath(), debian.RenderUnit(p.ExecutableName))
}
