package packager

import (
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/debian/constant"
)

func (p *Packager) SavePostRemove() {
	p.saveScript(constant.PostRemoveScript, debian.RenderPostRemove())
}
