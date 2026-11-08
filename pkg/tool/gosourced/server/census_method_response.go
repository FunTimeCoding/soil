package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func censusMethodResponse(m *result.CensusMethod) server.CensusMethod {
	callers := make([]server.CensusCaller, 0, len(m.Callers))

	for _, c := range m.Callers {
		callers = append(
			callers,
			server.CensusCaller{
				Kind:     c.Kind,
				Package:  c.Package,
				Function: c.Function,
				Count:    c.Count,
			},
		)
	}

	uses := make([]server.CensusUse, 0, len(m.Uses))

	for _, u := range m.Uses {
		uses = append(
			uses,
			server.CensusUse{Type: u.Type, Methods: listed(u.Methods)},
		)
	}

	return server.CensusMethod{
		Name:            m.Name,
		File:            m.File,
		Callers:         callers,
		Calls:           listed(m.Calls),
		Functions:       listed(m.Functions),
		Fields:          listed(m.Fields),
		Uses:            uses,
		Packages:        listed(m.Packages),
		Atomically:      m.Atomically,
		TransactionStep: m.TransactionStep,
		Tables: server.CensusTables{
			Read:    listed(m.Tables.Read),
			Written: listed(m.Tables.Written),
		},
		Reach: server.CensusTables{
			Read:    listed(m.Reach.Read),
			Written: listed(m.Reach.Written),
		},
	}
}
