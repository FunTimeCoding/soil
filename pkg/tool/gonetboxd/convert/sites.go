package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/site"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
)

func Sites(v []*site.Site) []*server.Site {
	result := make([]*server.Site, 0, len(v))

	for _, s := range v {
		result = append(result, Site(s))
	}

	return result
}
