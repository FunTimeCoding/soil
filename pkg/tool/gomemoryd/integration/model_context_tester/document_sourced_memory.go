package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
)

func (o *Tester) DocumentSourcedMemory() int64 {
	o.t.Helper()
	p := save_option.New()
	p.Name = "Retry"
	p.Content = "Document-sourced content."
	p.Description = ""
	p.Type = "reference"
	p.Scope = "alfa"
	p.ProvenanceFile = "canon/Example.yaml"
	p.ProvenanceAnchor = "Retry"
	identifier, e := o.Store().CreateMemory(p)
	assert.FatalOnError(o.t, e)

	return identifier
}
