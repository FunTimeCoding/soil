package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func createTagged(
	t *testing.T,
	o *service_tester.Tester,
	name string,
	tags ...string,
) *record.Memory {
	t.Helper()
	p := scopedOption(name, "")
	p.Type = "feedback"
	p.Tags = tags
	m, e := o.Service.CreateMemory(p)
	assert.FatalOnError(t, e)

	return m
}
