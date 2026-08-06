package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func AddressRange(r *internet_address_range.Range) *server.AddressRange {
	result := &server.AddressRange{
		Identifier: r.Identifier,
		Start:      r.Start,
		End:        r.End,
	}

	if r.Status != "" {
		result.Status = &r.Status
	}

	return result
}
