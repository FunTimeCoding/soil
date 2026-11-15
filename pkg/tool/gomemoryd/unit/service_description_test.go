package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/save_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func TestMultiLineDescriptionIsRefused(t *testing.T) {
	o := service_tester.New(t)
	p := save_tester.Scoped("alfa", "")
	p.Description = "one line</parameter>\n<parameter name=\"content\">body"
	_, e := o.Service.CreateMemory(p)
	assert.True(t, validation.Is(e))
	m := o.CreateTagged("bravo")
	q := save_tester.Scoped("bravo", "")
	q.Description = "still one line\nand a second"
	_, _, f := o.Service.UpdateMemory(m.Identifier, q)
	assert.True(t, validation.Is(f))
}
