package protocol

import "github.com/funtimecoding/soil/pkg/chromium/constant"

func (p *Protocol) Body() string {
	return p.Outer(constant.BodySelector)
}
