package service

import "github.com/funtimecoding/soil/pkg/netbox/internet_address_range"

func (s *Service) CreateInternetAddressRange(
	start string,
	end string,
	status string,
	description string,
) (*internet_address_range.Range, error) {
	return s.client.CreateInternetAddressRange(start, end, status, description)
}
