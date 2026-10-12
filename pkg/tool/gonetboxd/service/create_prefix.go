package service

import (
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/site"
)

func (s *Service) CreatePrefix(
	cidr string,
	siteName string,
	description string,
) (*prefix.Prefix, error) {
	var i *site.Site

	if siteName != "" {
		var e error
		i, e = s.client.SiteByName(siteName)

		if e != nil {
			return nil, e
		}
	}

	return s.client.CreatePrefix(cidr, i, description)
}
