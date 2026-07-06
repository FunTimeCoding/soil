package record

import (
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func New(
	v *hcloud.ZoneRRSet,
	z *zone.Zone,
) *Record {
	var values []string

	for _, r := range v.Records {
		values = append(values, r.Value)
	}

	return &Record{
		Name:     v.Name,
		Type:     string(v.Type),
		TTL:      v.TTL,
		Values:   values,
		ZoneName: z.Name,
		Raw:      v,
	}
}
