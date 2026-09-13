package layout

import "github.com/funtimecoding/soil/pkg/web/constant"

func (p *Page) extendedSource() string {
	if p.extended != "" {
		return p.extended
	}

	return constant.Extended
}
