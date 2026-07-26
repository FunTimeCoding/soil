package page

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Page) Format(f *option.Format) string {
	s := status.New(f).String(p.Name)

	if f.HasTag(constant.TagDense) {
		s.DetailLink(p.TinyLink, "Confluence", "")
	} else {
		s.DetailLink(p.Link, "Confluence", "")
	}

	return s.Format()
}
