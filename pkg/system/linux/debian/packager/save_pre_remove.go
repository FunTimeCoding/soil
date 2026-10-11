package packager

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian"
)

func (p *Packager) SavePreRemove() {
	p.saveScript(
		constant.DebianPreRemoveScript,
		debian.RenderPreRemove(p.ExecutableName),
	)
}
