package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/save_tester"
)

func (o *Tester) CreateBased() *record.Memory {
	o.t.Helper()
	p := save_tester.Scoped("based entry", "")
	p.Base = new("doc")
	p.Metadata = map[string]string{"kind": "mechanism"}
	m, e := o.Service().CreateMemory(p)
	assert.FatalOnError(o.t, e)

	return m
}
