package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
)

func convertFacet(f search.Facet) *server.Facet {
	result := &server.Facet{Key: f.Key, Distinct: f.Distinct}

	if len(f.Values) > 0 {
		result.Values = &f.Values
	}

	return result
}
