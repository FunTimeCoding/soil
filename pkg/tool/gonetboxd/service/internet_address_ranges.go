package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address_range"

func (s *Service) InternetAddressRanges() ([]*internet_address_range.Range, error) {
	return s.client.InternetAddressRanges()
}
