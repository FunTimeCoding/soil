package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"testing"
)

func createBased(
	t *testing.T,
	o *service_tester.Tester,
) *record.Memory {
	t.Helper()
	p := scopedOption("based entry", "")
	p.Base = new("doc")
	p.Metadata = map[string]string{"kind": "mechanism"}
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)

	return m
}
