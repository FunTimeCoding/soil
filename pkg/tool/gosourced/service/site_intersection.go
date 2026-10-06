package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/literal"
	"slices"
)

func siteIntersection(sites []*literal.Site) []string {
	if len(sites) == 0 {
		return nil
	}

	result := slices.Clone(sites[0].Fields)

	for _, site := range sites[1:] {
		result = slices.DeleteFunc(
			result,
			func(name string) bool { return !slices.Contains(site.Fields, name) },
		)
	}

	return result
}
