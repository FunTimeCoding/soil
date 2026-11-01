package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/report"
)

func diffResponse(r *report.Report) server.DiffResponse {
	result := server.DiffResponse{
		Counts:       []server.ClassCount{},
		Entries:      []server.Entry{},
		FromUnparsed: append([]string{}, r.FromUnparsed...),
		ToUnparsed:   append([]string{}, r.ToUnparsed...),
		Modules:      []server.ModuleRange{},
	}

	for _, m := range r.Modules {
		result.Modules = append(
			result.Modules,
			server.ModuleRange{Name: m.Name, From: m.From, To: m.To},
		)
	}

	for _, class := range constant.Classes {
		result.Counts = append(
			result.Counts,
			server.ClassCount{Class: class, Count: r.Count(class)},
		)
	}

	for _, e := range r.Entries {
		if e.Class == constant.ClassUnchanged {
			continue
		}

		result.Entries = append(
			result.Entries,
			server.Entry{
				Class: e.Class,
				From:  location(e.From),
				To:    location(e.To),
			},
		)
	}

	return result
}
