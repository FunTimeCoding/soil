package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"time"
)

func (s *Server) DocumentSourcedMemory() int64 {
	s.t.Helper()
	p := save_option.New()
	p.Name = "Retry"
	p.Content = "Document-sourced content."
	p.Description = ""
	p.Type = "reference"
	p.Scope = "alfa"
	p.ProvenanceFile = "canon/Example.yaml"
	p.ProvenanceAnchor = "Retry"
	identifier, e := s.store.CreateMemory(p, time.Now())
	assert.FatalOnError(s.t, e)

	return identifier
}
