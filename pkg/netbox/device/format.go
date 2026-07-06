package device

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (d *Device) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(d.Identifier)
	}

	s.String(
		d.formatName(f),
		d.PrimaryAddress,
		d.formatSerial(f),
	).RawList(d.Raw)
	s.DetailLink(d.Link, "NetBox", "")

	if t := d.formatTags(f); t != "" {
		s.Line("  Tags: %s", t)
	}

	if v := d.formatComment(f); v != "" {
		s.Line("  Comment: %s", v)
	}

	return s.Format()
}
