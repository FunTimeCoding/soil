package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func AddressRanges(v []*internet_address_range.Range) []*server.AddressRange {
	result := make([]*server.AddressRange, 0, len(v))

	for _, r := range v {
		result = append(result, AddressRange(r))
	}

	return result
}
