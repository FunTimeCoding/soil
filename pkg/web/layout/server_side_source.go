package layout

import "github.com/funtimecoding/soil/pkg/web/constant"

func (p *Page) serverSideSource() string {
	if p.serverSide != "" {
		return p.serverSide
	}

	return constant.ServerSide
}
