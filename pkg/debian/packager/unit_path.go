package packager

import (
	"github.com/funtimecoding/soil/pkg/debian/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func (p *Packager) unitPath() string {
	return join.Absolute(
		p.UnitRoot,
		key_value.Dot(p.ExecutableName, constant.ServiceExtension),
	)
}
