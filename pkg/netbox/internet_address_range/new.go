package internet_address_range

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.IPRange) *Range {
	s := v.GetStatus()

	return &Range{
		Identifier: v.GetId(),
		Start:      v.GetStartAddress(),
		End:        v.GetEndAddress(),
		Status:     string(s.GetValue()),
		Raw:        v,
	}
}
