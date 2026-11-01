package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/intent"
)

func convertPlan(p *intent.Plan) server.ReconcilePlan {
	result := server.ReconcilePlan{
		Hypervisor: p.Hypervisor,
		Written:    p.Written,
		Changes:    []server.ReconcileChange{},
		Unknown:    []string{},
		Ambiguous:  []string{},
	}

	for _, c := range p.Changes {
		result.Changes = append(
			result.Changes,
			server.ReconcileChange{
				Name: c.Name,
				Node: c.Node,
				From: c.From,
				To:   c.To,
			},
		)
	}

	result.Unknown = append(result.Unknown, p.Unknown...)
	result.Ambiguous = append(result.Ambiguous, p.Ambiguous...)

	return result
}
