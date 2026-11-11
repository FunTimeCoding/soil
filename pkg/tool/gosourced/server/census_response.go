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

	unloaded := make([]server.CensusUnloaded, 0, len(c.Unloaded))

	for _, u := range c.Unloaded {
		unloaded = append(
			unloaded,
			server.CensusUnloaded{
				Package: u.Package,
				Cause:   u.Cause,
				Lost:    u.Lost,
			},
		)
	}

	return server.Census{
		Package:  c.Package,
		Type:     c.Type,
		Fields:   listed(c.Fields),
		Methods:  methods,
		Unloaded: unloaded,
	}
}
