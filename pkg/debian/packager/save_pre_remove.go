package packager

import (
	"github.com/funtimecoding/soil/pkg/debian"
	"github.com/funtimecoding/soil/pkg/debian/constant"
)

func (p *Packager) SavePreRemove() {
	p.saveScript(
		constant.PreRemoveScript,
		debian.RenderPreRemove(p.ExecutableName),
	)
}
