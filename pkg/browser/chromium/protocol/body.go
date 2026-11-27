package protocol

import "github.com/funtimecoding/soil/pkg/browser/constant"

func (p *Protocol) Body() (string, error) {
	return p.Outer(constant.ChromiumBodySelector)
}
