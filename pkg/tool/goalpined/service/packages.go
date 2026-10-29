package service

import "github.com/funtimecoding/soil/pkg/system/linux/alpine/types/listing"

func (s *Service) Packages(name string) ([]*listing.Listing, error) {
	result, e := listing.Indexes(s.packageRoot)

	if e != nil {
		return nil, e
	}

	return listing.Filter(result, name), nil
}
