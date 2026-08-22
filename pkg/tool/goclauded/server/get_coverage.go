package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"time"
)

func (s *Server) GetCoverage(
	_ context.Context,
	_ server.GetCoverageRequestObject,
) (server.GetCoverageResponseObject, error) {
	var servers []server.CoverageServer

	for _, e := range s.service.Coverage() {
		entry := server.CoverageServer{
			Name:        e.Name,
			Configured:  e.Configured,
			Registered:  e.Registered,
			UsedTotal:   e.UsedTotal,
			UsedRecent:  e.UsedRecent,
			CallsTotal:  e.CallsTotal,
			CallsRecent: e.CallsRecent,
		}

		if e.Path != "" {
			entry.Path = new(e.Path)
		}

		if !e.LastUsed.IsZero() {
			entry.LastUsed = new(e.LastUsed.Format(time.RFC3339))
		}

		for _, t := range e.Tools {
			tool := server.CoverageTool{
				Name:        t.Name,
				Registered:  t.Registered,
				CallsTotal:  t.CallsTotal,
				CallsRecent: t.CallsRecent,
			}

			if !t.LastUsed.IsZero() {
				tool.LastUsed = new(t.LastUsed.Format(time.RFC3339))
			}

			entry.Tools = append(entry.Tools, tool)
		}

		servers = append(servers, entry)
	}

	return server.GetCoverage200JSONResponse{Servers: servers}, nil
}
