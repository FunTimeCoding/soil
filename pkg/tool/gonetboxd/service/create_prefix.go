package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/prefix"
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreatePrefix(
	cidr string,
	siteName string,
	description string,
) (*prefix.Prefix, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*prefix.Prefix, error) {
			var i *site.Site

			if siteName != "" {
				var e error
				i, e = s.client.Site(siteName)

				if e != nil {
					return nil, e
				}
			}

			return s.client.CreatePrefix(cidr, i, description)
		},
	)
}
