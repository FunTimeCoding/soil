package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"testing"
)

func createNamed(
	t *testing.T,
	o *service_tester.Tester,
	name string,
	scope string,
	content string,
) *record.Memory {
	t.Helper()
	p := scopedOption(name, scope)
	p.Content = content
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)

	return m
}
