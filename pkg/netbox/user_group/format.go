package user_group

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (g *Group) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		s.Integer32(g.Identifier)
	}

	s.String(g.formatName(f)).RawList(g.Raw)

	return s.Format()
}
