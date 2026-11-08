package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func censusResponse(c *result.Census) server.Census {
	methods := make([]server.CensusMethod, 0, len(c.Methods))

	for _, m := range c.Methods {
		methods = append(methods, censusMethodResponse(m))
	}

	return server.Census{
		Package: c.Package,
		Type:    c.Type,
		Fields:  listed(c.Fields),
		Methods: methods,
	}
}
