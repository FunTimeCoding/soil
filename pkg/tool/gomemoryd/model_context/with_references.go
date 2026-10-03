package model_context

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
)

func (s *Server) withReferences(
	text string,
	m *store.Memory,
) string {
	findings := s.service.References(m)

	if len(findings) == 0 {
		return text
	}

	return join.NewLine(
		append(
			[]string{text, constant.ReferenceHeader},
			reference.Lines(findings)...,
		),
	)
}
