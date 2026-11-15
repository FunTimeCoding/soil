package service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/save_tester"
)

func (o *Tester) CreateNamed(
	name string,
	scope string,
	content string,
) *record.Memory {
	o.t.Helper()
	p := save_tester.Scoped(name, scope)
	p.Content = content
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(o.t, e)

	return m
}
