package alert

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (a *Alert) Format(f *option.Format) string {
	s := status.New(f)

	if v := a.formatEntity(f); v != "" {
		s.String(v)
	}

	if f.HasTag(tag.Category) {
		if v := a.formatCategory(f); v != "" {
			s.String(v)
		}
	}

	if a.Entity == "" && a.Category == "" {
		s.String(a.formatName(f))
	}

	s.String(a.formatPriority(f))

	if !f.HasTag(tag.Filter) && !f.HasTag(tag.Dense) {
		s.String(a.formatStatus(f))
	}

	s.String(a.formatTeamName(f), a.formatAge())

	if o := a.formatOwner(f); o != "" {
		s.String(o)
	}

	if l := a.flags(f); len(l) > 0 {
		s.String(join.Comma(l))
	}

	if f.ShowExtended {
		a.extended(s, f)
	}

	if a.RawList != nil {
		s.RawList(a.RawList)
	}

	if a.RawDetail != nil {
		s.RawDetail(a.RawDetail)
	}

	return s.Format()
}
