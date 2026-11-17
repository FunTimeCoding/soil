package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"time"
)

func (s *Server) EmptyContentParent() int64 {
	s.t.Helper()
	p := save_option.New()
	p.Name = "Document"
	p.Content = ""
	p.Description = "document description"
	p.Type = "reference"
	p.ProvenanceFile = "canon/Document.yaml"
	identifier, e := s.store.CreateMemory(p, time.Now())
	assert.FatalOnError(s.t, e)

	return identifier
}
