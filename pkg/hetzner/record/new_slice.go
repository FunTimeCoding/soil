package record

import (
	"github.com/funtimecoding/soil/pkg/hetzner/zone"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func NewSlice(
	v []*hcloud.ZoneRRSet,
	z *zone.Zone,
) []*Record {
	var result []*Record

	for _, e := range v {
		result = append(result, New(e, z))
	}

	return result
}
