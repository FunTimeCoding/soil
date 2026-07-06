package rack

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
)

func (r *Rack) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(r.Identifier)
	}

	s.String(r.formatName(f)).RawList(r.Raw)
	s.DetailLink(r.Link, "NetBox", "")

	if tags := r.tags(); len(tags) > 0 {
		s.Line("  Tags: %s", join.Comma(tags))
	}

	s.Line("  Status: %s", r.status(f))
	s.Line("  Site: %s", r.Raw.Site.GetName())

	if r.Raw.Role.IsSet() {
		s.Line("  Role: %s", r.Raw.Role.Get().GetName())
	}

	if r.Raw.Location.IsSet() {
		s.Line("  Location: %s", r.Raw.Location.Get().GetName())
	}

	if r.Raw.Tenant.IsSet() {
		s.Line("  Tenant: %s", r.Raw.Tenant.Get().GetName())
	}

	s.Line(
		"  Created: %s",
		r.Raw.Created.Get().Format(time.DateMinute),
	)
	s.Line(
		"  Updated: %s",
		r.Raw.LastUpdated.Get().Format(time.DateMinute),
	)
	s.Line("  Devices: %d", r.Raw.DeviceCount)
	s.Line("  Unit height: %d", r.Raw.UHeight)

	return s.Format()
}
