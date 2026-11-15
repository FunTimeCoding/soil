package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
)

func usageMethodResponse(m *result.UsageMethod) server.UsageMethod {
	callers := make([]server.UsageCaller, 0, len(m.Callers))

	for _, c := range m.Callers {
		callers = append(
			callers,
			server.UsageCaller{
				Kind:     c.Kind,
				Package:  c.Package,
				Function: c.Function,
				Count:    c.Count,
			},
		)
	}

	uses := make([]server.UsedType, 0, len(m.Uses))

	for _, u := range m.Uses {
		uses = append(
			uses,
			server.UsedType{Type: u.Type, Methods: listed(u.Methods)},
		)
	}

	return server.UsageMethod{
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
		Tables: server.UsageTables{
			Read:    listed(m.Tables.Read),
			Written: listed(m.Tables.Written),
		},
		Reach: server.UsageTables{
			Read:    listed(m.Reach.Read),
			Written: listed(m.Reach.Written),
		},
	}
}
