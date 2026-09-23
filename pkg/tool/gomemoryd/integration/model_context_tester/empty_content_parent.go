package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (o *Tester) EmptyContentParent() int64 {
	o.t.Helper()
	p := save_option.New()
	p.Name = "Document"
	p.Content = ""
	p.Description = "document description"
	p.Type = "reference"
	p.ProvenanceFile = "canon/Document.yaml"
	identifier, e := o.Store().CreateMemory(p)
	assert.FatalOnError(o.t, e)

	return identifier
}
