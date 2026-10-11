package packager

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian"
)

func (p *Packager) SavePostRemove() {
	p.saveScript(constant.DebianPostRemoveScript, debian.RenderPostRemove())
}
