package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func usageResponse(c *result.Usage) server.Usage {
	methods := make([]server.UsageMethod, 0, len(c.Methods))

	for _, m := range c.Methods {
		methods = append(methods, usageMethodResponse(m))
	}

	unloaded := make([]server.UsageUnloaded, 0, len(c.Unloaded))

	for _, u := range c.Unloaded {
		unloaded = append(
			unloaded,
			server.UsageUnloaded{
				Package: u.Package,
				Cause:   u.Cause,
				Lost:    u.Lost,
			},
		)
	}

	return server.Usage{
		Package:  c.Package,
		Type:     c.Type,
		Fields:   listed(c.Fields),
		Methods:  methods,
		Unloaded: unloaded,
	}
}
